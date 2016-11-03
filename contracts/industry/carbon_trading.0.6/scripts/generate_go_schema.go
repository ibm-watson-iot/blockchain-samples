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

type OBJ map[string]interface{}

// contents of "generate.json" colocated in scripts folder with this script
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
        if objName == "definitions" {
            objName = "#/" + objName
        } else {
            objName = "#/definitions/" + objName
        }
    } 
    s := strings.Split(objName, "/")
    var found bool
    // crawl the levels, skipping the # root
    for i := 1; i < len(s); i++ {
        schema, found = (schema[s[i]]).(map[string]interface{})
        if !found {
            fmt.Println("** ERR ** illegal selector for field: ", objName)
            return nil
        }
    }
    return schema
}

// replaces all references recursively in the passed-in object (subschema) using the passed-in schema
func replaceReferences (schema map[string]interface{}, obj map[string]interface{}) (map[string]interface{}) {
    for k, v := range obj {
        switch /* t := */ v.(type) { 
            default:
                //fmt.Printf("k: [%s] type: [%s]\n", k, t)
            case map[string]interface{}:
                //fmt.Printf("k: [%s] is a map of size: %d\n", k, len(k))
                r, found := v.(map[string]interface{})["$ref"]
                if found {
                    // it is a reference so replace it and recursively replace references
                    obj[k] = replaceReferences(schema, getObject(schema, r.(string)))
                } else {
                    obj[k] = replaceReferences(schema, v.(map[string]interface{}))
                }
        } 
    }
    return obj 
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
    var schemas map[string]interface{} = make(map[string]interface{})
    var outString string = "package main\n\nvar schemas = `\n"

    var filename string = config.Schemas.GoSchemaFilename
    var apiFunctions []string = config.Schemas.API
    var elementNames []string = config.Schemas.GoSchemaElements
    
    var functionKey string = "API"
    var objectModelKey string = "objectModelSchemas"
    
    schemas[functionKey] = interface{}(make(map[string]interface{}))
    schemas[objectModelKey] = interface{}(make(map[string]interface{}))

    fmt.Printf("Generate Go SCHEMA file %s for: \n   %s and: \n   %s\n", filename, apiFunctions, elementNames)

    // grab the event API functions for input
    for i := range apiFunctions {
        functionSchemaName := "API/properties/" + apiFunctions[i]
        functionName := apiFunctions[i]
        obj = getObject(schema, functionSchemaName)
        if obj == nil {
            fmt.Printf("** Warning** %s returned nil from getObject\n", functionSchemaName)
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
            fmt.Printf("** Warning** %s returned nil from getObject\n", elementName)
            return
        }
        schemas[objectModelKey].(map[string]interface{})[elementName] = obj 
    }
    
    // marshal for output to file     
    schemaOut, err := json.MarshalIndent(&schemas, "", "    ")
    if err != nil {
        fmt.Printf("** Warning** cannot marshal schema file output for writing\n")
        return
    }
    outString += string(schemaOut) + "`"
    ioutil.WriteFile(filename, []byte(outString), 0644)
}

func sampleType(obj map[string]interface{}, elementName string) (interface{}) {
    t, found := obj["type"].(string)
    if !found {
        fmt.Printf("**Warning** Element %s has no type field\n", elementName)
        return "NO TYPE PROPERTY"
    }
    switch t {
        case "number" :
            return 123.456
        case "integer" :
            return 789
        case "string" :
            if strings.ToLower(elementName) == "timestamp" {
                return time.Now().Format(time.RFC3339Nano)
            }
            example, found := obj["example"].(string)
            if found && len(example) > 0 {
                return example
            }
            def, found := obj["default"].(string)
            if found && len(def) > 0 {
                return def
            }
            enum, found := obj["enum"].([]interface{})
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
            desc, found := obj["description"].(string)
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
            var items, found = obj["items"].(map[string]interface{})
            if (!found) {
                fmt.Printf("**Warning** Element %s is array with no items property\n", elementName)
                return "ARRAY WITH NO ITEMS PROPERTY"
            }
            return arrayFromSchema(items, elementName)
        case "object" : {
            props, found := obj["properties"]
            if !found {
                fmt.Printf("** Warning** %s is type object yet has no properties in SampleType\n", elementName)
                return "INVALID OBJECT - MISSING PROPERTIES"
            }
            objOut := make(map[string]interface{})
            for k, v := range props.(map[string]interface{}) {
                //fmt.Printf("Visiting key %s with value %s\n", k, v)
                if v == nil {
                    fmt.Printf("** Warning** Key %s has NIL value in SampleType\n", k)
                    return "INVALID OBJECT - " + fmt.Sprintf("Key %s has NIL value in SampleType\n", k)
                }
                objOut[k] = sampleType(v.(map[string]interface{}), k)
            }
            return objOut
        }
    }
    fmt.Printf("** Warning** UNKNOWN TYPE in SampleType: %s\n", t)
    return fmt.Sprintf("UNKNOWN TYPE in SampleType: %s\n", t)
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
    var samples map[string]interface{} = make(map[string]interface{})
    var outString string = "package main\n\nvar samples = `\n"

    var filename string = config.Samples.GoSampleFilename
    var elementNames []string = config.Samples.GoSampleElements

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
                fmt.Printf("** Warning** %s returned nil from getObject\n", elementName)
                return
            }
        }
        samples[elementName] = sampleType(obj, elementName) 
    }
    samplesOut, err := json.MarshalIndent(&samples, "", "    ")
    if err != nil {
        fmt.Println("** Warning** cannot marshal sample file output for writing")
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
    var line int = 1
    var lineOut int = 1    
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
        fmt.Printf("** ERR [%s] opening schema file at %s\n", err, config.Schemas.SchemaFilename)
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
    schema = replaceReferences(schema, schema)
    
    // generate the Go files that the contract needs -- for now, complete schema and 
    // event schema and sample object 
    
    generateGoSchemaFile(schema, config)
    generateGoSampleFile(schema, config)
    
    // experimental
    //generateGoObjectModel(schema, config)
    
    // TODO generate js object model?? Java??
        
}