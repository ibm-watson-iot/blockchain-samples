/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Kim Letkeman - Initial Contribution
*/


// IoT Blockchain Demo Smart Contract
// KL 03 Mar 2016 Generate schema and event subschema Go files for contract v3.1
// KL 04-07 Mar 2016 testing of schema, adaptation of output to contract 3.0.2,
//                   addition of config file generate.yaml
// KL 13 Mar 2016 Changed from yaml (lesser GPL) to JSON for config 
// KL 8 June 2016 Supporting complex events and the "oneOf" keyword, better support
//                for arrays, cleanup lint issues

package main

import (
    "bufio"
    "io/ioutil"
    "os"
    "strings"
    "fmt"
	"encoding/json"
    "time"
    "path/filepath"
)

// Config defines contents of "generate.json" colocated in scripts folder with this script
type Config struct {
    Schemas struct {
        SchemaFilename string           `json:"schemaFilename"`
        GoSchemaFilename string         `json:"goSchemaFilename"`
        GoSchemaElements []string       `json:"goSchemaElements"`
        API []string                    `json:"API"`
    } `json:"schemas"`
    Samples struct {
        GoSampleFilename string         `json:"goSampleFilename"`
        GoSampleElements []string       `json:"goSampleElements"`
    } `json:"samples"`
    ObjectModels struct {
        ObjectModelElements []string    `json:"generateGoObjectsFrom"`
    } `json:"objectModels"`
}

// can print very accurate syntax errors as found by the JSON marshaler
// relies on the offset table created when reading the schema JSON file and expunging 
// comments and blank lines
func printSyntaxError(js string, off *[5000]int, err interface{}) {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
        fmt.Println("*********** ERR trying to get syntax error location **************\n", err)
		return
	}
	
	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}
	
	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset) - start -1
	
	fmt.Printf("Error in line %d: %s \n", off[line]+1, err)
	fmt.Printf("%s\n%s^\n\n", js[start:end], strings.Repeat(" ", pos))
}

// retrieves a subschema object via the reference path; handles root node references and 
// references starting after definitions; does not handle external file references yet
func getObject (schema map[string]interface{}, objName string) (map[string]interface{}) {
    // return a copy of the selected object
    // handles full path, or path starting after definitions
    if !strings.HasPrefix(objName, "#/definitions/") {
        objName = "#/definitions/" + objName
    } 
    s := strings.Split(objName, "/")
    // crawl the levels, skipping the # root
    for i := 1; i < len(s); i++ {
        props, found := (schema["properties"]).(map[string]interface{})
        if found {
            schema, found = (props[s[i]]).(map[string]interface{})
        } else {
            schema, found = (schema[s[i]]).(map[string]interface{})
        }
        if !found {
            fmt.Printf("schema[s[i]] called %s looks like: %+v\n", objName, schema[s[i]])
            fmt.Printf("** ERR ** getObject illegal selector %s at level %d called %s\n", objName, i, s[i])
            return nil
        }
    }
    return schema
}

// replaces all references recursively in the passed-in object (subschema) using the passed-in schema
func replaceReferences (schema map[string]interface{}, obj interface{}) (interface{}) {
    oArr, isArr := obj.([]interface{})
    oMap, isMap := obj.(map[string]interface{})
    switch {
        default:
            return obj
        case isArr:
            //fmt.Printf("ARR [%s:%+v]\n", k, v)
            for k, v := range oArr {
                r, found := v.(map[string]interface{})
                if found {
                    ref, found := r["$ref"] 
                    if found {
                        // it is a reference so replace it and recursively replace references
                        oArr[k] = replaceReferences(schema, getObject(schema, ref.(string)))
                    } else {
                        oArr[k] = replaceReferences(schema, v)
                    }
                } else {
                    //fmt.Printf("** WARN ** array member not a map object [%d:%+v]\n", k, v)
                }
            }
            return oArr 
        case isMap:
            //fmt.Printf("MAP [%s:%+v]\n", k, v)
            for k, v := range oMap {
                if k == "$ref" {
                    // it is a reference so replace it and recursively replace references
                    //fmt.Printf("** INFO ** Should be $ref [%s:%+v]\n", k, v)
                    oMap = replaceReferences(schema, getObject(schema, v.(string))).(map[string]interface{})
                } else {
                    oMap[k] = replaceReferences(schema, v)
                }
            }
            return oMap
        } 
}

// If a reference exists at any level in the passed-in schema, this will return true
// Recurses through every level of the map
func referencesExist (schema map[string]interface{}) (bool) {
    _, exists := schema["$ref"]
    if exists {
        return true
    } 
    for _, v := range schema {
        switch v.(type) { 
            case map[string]interface{}:
                if referencesExist(v.(map[string]interface{})) {
                    return true
                }
        }
    }
    return false 
}

// Generates a file <munged elementName>.go to contain a string literal for the pretty version 
// of the schema with all references resolved. In the same file, creates a sample JSON that
// can be used to show a complete structure of the object.  
func generateGoSchemaFile(schema map[string]interface{}, config Config) {
    var obj map[string]interface{}
    var schemas = make(map[string]interface{})
    var outString = "package main\n\nvar schemas = `\n"

    var filename = config.Schemas.GoSchemaFilename
    var apiFunctions = config.Schemas.API
    var elementNames = config.Schemas.GoSchemaElements
    
    var functionKey = "API"
    var objectModelKey = "objectModelSchemas"
    
    schemas[functionKey] = interface{}(make(map[string]interface{}))
    schemas[objectModelKey] = interface{}(make(map[string]interface{}))

    fmt.Printf("Generate Go SCHEMA file %s for: \n   %s and: \n   %s\n", filename, apiFunctions, elementNames)

    // grab the event API functions for input
    for i := range apiFunctions {
        functionSchemaName := "#/definitions/API/" + apiFunctions[i]
        functionName := apiFunctions[i]
        obj = getObject(schema, functionSchemaName)
        if obj == nil {
            fmt.Printf("** WARN ** %s returned nil from getObject\n", functionSchemaName)
            return
        }
        schemas[functionKey].(map[string]interface{})[functionName] = obj 
    }     

    // grab the elements requested (these are useful separately even though
    // they obviously appear already as part of the event API functions)
    for i := range elementNames {
        elementName := elementNames[i]
        obj = getObject(schema, elementName)
        if obj == nil {
            fmt.Printf("** ERR ** %s returned nil from getObject\n", elementName)
            return
        }
        schemas[objectModelKey].(map[string]interface{})[elementName] = obj 
    }
    
    // marshal for output to file     
    schemaOut, err := json.MarshalIndent(&schemas, "", "    ")
    if err != nil {
        fmt.Printf("** ERR ** cannot marshal schema file output for writing\n")
        return
    }
    outString += string(schemaOut) + "`"
    ioutil.WriteFile(filename, []byte(outString), 0644)
}

func sampleType(obj interface{}, elementName string) (interface{}) {
    o, found := obj.(map[string]interface{})
    if (!found) {
        return "SCHEMA ELEMENT " + elementName + " IS NOT MAP"
    }
    t, found := o["type"].(string)
    if !found {
        //fmt.Printf("** WARN ** Element %s has no type field\n", elementName)
        //fmt.Printf("Element missing type is: %s [%v]\n\n", elementName, o)
        if elementName == "oneOf" {
            return o
        }
        return "NO TYPE PROPERTY"
    }
    switch t {
        default :
            fmt.Printf("** WARN ** Unknown type in sampleType %s\n", t)
        case "number" :
            return 123.456
        case "integer" :
            return 789
        case "string" :
            if strings.ToLower(elementName) == "timestamp" {
                return time.Now().Format(time.RFC3339Nano)
            }
            example, found := o["example"].(string)
            if found && len(example) > 0 {
                return example
            }
            def, found := o["default"].(string)
            if found && len(def) > 0 {
                return def
            }
            enum, found := o["enum"].([]interface{})
            if found {
                if len(enum) > 1 {
                    return enum[1]
                } 
                if len(enum) > 0 { 
                    return enum[0]
                }
            }
            // description is a good alternate choice for sample data since it
            // explains the prospective contents
            desc, found := o["description"].(string)
            if found && len(desc) > 0 {
                return desc
            }
            // if nothing else ...
            return "carpe noctem"
        case "null" :
            return nil
        case "boolean" :
            return true
        case "array" :
            var items, found = o["items"].(map[string]interface{})
            if (!found) {
                fmt.Printf("** WARN ** Element %s is array with no items property\n", elementName)
                return "ARRAY WITH NO ITEMS PROPERTY"
            }
            return arrayFromSchema(items, elementName)
        case "object" : {
            props, found := o["properties"]
            if !found {
                fmt.Printf("** WARN ** %s is type object yet has no properties in SampleType\n", elementName)
                return "INVALID OBJECT - MISSING PROPERTIES"
            }
            objOut := make(map[string]interface{})
            for k, v := range props.(map[string]interface{}) {
                //fmt.Printf("Visiting key %s with value %s\n", k, v)
                if v == nil {
                    fmt.Printf("** WARN ** Key %s has NIL value in SampleType\n", k)
                    return "INVALID OBJECT - " + fmt.Sprintf("Key %s has NIL value in SampleType\n", k)
                }
                aArr, isArr := v.([]interface{})
                aMap, isMap := v.(map[string]interface{})
                switch {
                    case isArr:
                        if "oneOf" == k {
                            aOut := make([]interface{}, len(aArr))
                            // outer loop is anonymous objects
                            for k2, v2 := range aArr {
                                //fmt.Printf("SAMTYP outer OneOf: %d [%v]\n", k2, v2)
                                vObj, found := v2.(map[string]interface{})
                                if found {
                                    // inner loop should find one named object
                                    for k3, v3 := range vObj {
                                        tmp := make(map[string]interface{}, 1)
                                        //fmt.Printf("SAMTYP inner OneOf: %s [%v]\n", k3, v3)
                                        //printObject(k3, v3)
                                        tmp[k3] = sampleType(v3, k3) 
                                        aOut[k2] = tmp
                                    }
                                }
                                objOut[k] = aOut
                            }
                        } else {
                            objOut[k] = "UNKNOWN ARRAY OBJECT"
                        }
                    case isMap:
                        objOut[k] = sampleType(aMap, k)
                }
            }
            return objOut
        }
    }
    fmt.Printf("** WARN ** UNKNOWN TYPE in SampleType: %s\n", t)
    return fmt.Sprintf("UNKNOWN TYPE in SampleType: %s\n", t)
}

func printObject(elementName string, obj interface{}) {
    aMap, isMap := obj.(map[string]interface{})
    aArr, isArr := obj.([]interface{})
    switch {
        case isArr:
            fmt.Printf("Element: %s is an ARRAY\n", elementName)
            for k, v := range aArr {
                fmt.Printf("[%d] : %+v\n\n", k, v)
            }
        case isMap:
            fmt.Printf("Element: %s is a MAP\n", elementName)
            for k, v := range aMap {
                fmt.Printf("[%s] : %+v\n\n", k, v)
            }
        default:
            fmt.Printf("Element: %s is of UNKNOWN shape\n", elementName)
    }
}

// Generate a sample array from a schema 
func arrayFromSchema(schema map[string]interface{}, elementName string) (interface{}) {
    enum, found := schema["enum"]
    if found {
        // there is a set of enums, just use it
        return enum
    }
    return []interface{}{sampleType(schema, elementName)}
}

// Generates a file <munged elementName>.go to contain a string literal for the pretty version 
// of the schema with all references resolved. In the same file, creates a sample JSON that
// can be used to show a complete structure of the object.  
func generateGoSampleFile(schema map[string]interface{}, config Config) {
    var obj map[string]interface{}
    var samples = make(map[string]interface{})
    var outString = "package main\n\nvar samples = `\n"

    var filename = config.Samples.GoSampleFilename
    var elementNames = config.Samples.GoSampleElements

    fmt.Printf("Generate Go SAMPLE file %s for: \n   %s\n", filename, elementNames)

    for i := range elementNames {
        elementName := elementNames[i]
        if elementName == "schema" {
            // sample of the entire schema, can it even work?
            obj = schema
        } else {
            // use the schema subset
            obj = getObject(schema, elementName)
            if obj == nil {
                fmt.Printf("** WARN ** %s returned nil from getObject\n", elementName)
                return
            }
        }
        samples[elementName] = sampleType(obj, elementName) 
    }
    samplesOut, err := json.MarshalIndent(&samples, "", "    ")
    if err != nil {
        fmt.Println("** ERR ** cannot marshal sample file output for writing")
        return
    }
    outString += string(samplesOut) + "`"
    ioutil.WriteFile(filename, []byte(outString), 0644)
}

func generateGoObjectModel(schema map[string]interface{}, config Config) () {
    for i := range config.ObjectModels.ObjectModelElements {
        fmt.Println("Generating object model for: ", 
                    config.ObjectModels.ObjectModelElements[i])
        obj := getObject(schema, config.ObjectModels.ObjectModelElements[i])
        fmt.Printf("%s: %s\n\n", config.ObjectModels.ObjectModelElements[i], obj)
    }
}

// Reads payloadschema.json api file
// encodes as a string literal in payloadschema.go
func main() {

    var configFileName = "generate.json"

    var api string
    var line = 1
    var lineOut = 1
    var offsets [5000]int
    
    // read the configuration from the json file
    filename, _ := filepath.Abs("./scripts/" + configFileName)
    fmt.Printf("JSON CONFIG FILEPATH:\n   %s\n", filename)
    jsonFile, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("error reading json file")
        panic(err)
    }
    var config Config
    err = json.Unmarshal(jsonFile, &config)
    if err != nil {
        fmt.Println("error unmarshaling json config")
        panic(err)
    }

    // read the schema file, stripping comments and blank lines, calculate offsets for error output
    file, err := os.Open(config.Schemas.SchemaFilename)
    if err != nil {
        fmt.Printf("** ERR ** [%s] opening schema file at %s\n", err, config.Schemas.SchemaFilename)
        return
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        ts := strings.TrimSpace(scanner.Text()) 
        if strings.HasPrefix(ts, "#") {
            fmt.Println("Line: ", line, " is a comment")
        } else if ts == "" {
            fmt.Println("Line: ", line, " is blank")
        } else {
            api += ts + "\n"
            lineOut++
        }
        offsets[lineOut] = line
        line++
    }
    
    // verify the JSON format by unmarshaling it into a map

    var schema map[string]interface{}
    err = json.Unmarshal([]byte(api), &schema)
    if err != nil {
        fmt.Println("*********** UNMARSHAL ERR **************\n", err)
        printSyntaxError(api, &offsets, err)
        return
    }
    
    // Looks tricky, but simply creates an output with references resolved 
    // from the schema, and another object and passes it back. I used to 
    // call it for each object, but much simpler to call it once for the 
    // whole schema and simply pick off the objects we want for subschemas
    // and samples.
    schema = replaceReferences(schema, schema).(map[string]interface{})
    
    // generate the Go files that the contract needs -- for now, complete schema and 
    // event schema and sample object 
    
    generateGoSchemaFile(schema, config)
    generateGoSampleFile(schema, config)
    
    // experimental
    //generateGoObjectModel(schema, config)
    
    // TODO generate js object model?? Java??
        
}