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

// KL 2016 Nov 20 rewrite, as funky behaviors had crept in with old algorithm, new algorithm uses lookup
//                table for references, which will improve accuracy and performance

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config defines contents of "generate.json" colocated in scripts folder with this script
type Config struct {
	Schemas struct {
		SchemaFilename   string   `json:"schemaFilename"`
		GoSchemaFilename string   `json:"goSchemaFilename"`
		API              []string `json:"API"`
		Model            []string `json:"Model"`
	} `json:"schemas"`
	Samples struct {
		GoSampleFilename string   `json:"goSampleFilename"`
		API              []string `json:"API"`
		Model            []string `json:"Model"`
	} `json:"samples"`
}

var configFile = flag.String("configFile", "generate.json", "json file that selects API to be exposed")
var verbose = flag.Bool("debug", false, "prints information during processing to help debug schema issues")
var config Config
var finalschema map[string]interface{}
var lookup = make(map[string]interface{}, 0)

// PrettyPrint returns an indented JSON stringified object
func PrettyPrint(m interface{}) string {
	bytes, _ := json.MarshalIndent(m, "", "    ")
	return string(bytes)
}

// PrettyPrintBytes returns an indented JSON stringified object as []bytes
func PrettyPrintBytes(m interface{}) []byte {
	bytes, _ := json.MarshalIndent(m, "", "    ")
	return bytes
}

// DeepMergeMap all levels of a src map into a dst map and return dst
func DeepMergeMap(srcIn map[string]interface{}, dstIn map[string]interface{}) map[string]interface{} {
	for k, v := range srcIn {
		switch v.(type) {
		case map[string]interface{}:
			dstv, found := dstIn[k]
			if found {
				// recursive DeepMerge into existing key
				dstIn[k] = DeepMergeMap(v.(map[string]interface{}), dstv.(map[string]interface{}))
			} else {
				// copy src to dst at same key
				dstIn[k] = v
			}
		default:
			// copy discrete type
			dstIn[k] = v
		}
	}
	return dstIn
}

// can print very accurate syntax errors as found by the JSON marshaler
// relies on the offset table created when reading the schema JSON file and expunging
// comments and blank lines
func printSyntaxErrorIncludes(js string, err interface{}) string {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
		fmt.Println("*********** ERR trying to get syntax error location **************\n", err)
		return "*********** ERR trying to get syntax error location **************"
	}

	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}

	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset)-start-1

	e := fmt.Sprintf("Error in line %d: %s \n", line, err)
	e += fmt.Sprintf("%s\n%s^\n\n", js[start:end], strings.Repeat(" ", pos))
	fmt.Println(e)
	return e
}

// can print very accurate syntax errors as found by the JSON marshaler
// relies on the offset table created when reading the schema JSON file and expunging
// comments and blank lines
func printSyntaxErrorOffsets(js string, off *[5000]int, err interface{}) string {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
		fmt.Println("*********** ERR trying to get syntax error location **************\n", err)
		return "*********** ERR trying to get syntax error location **************"
	}

	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}

	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset)-start-1

	e := fmt.Sprintf("Error in line %d: %s \n", off[line]+1, err)
	e += fmt.Sprintf("%s\n%s^\n\n", js[start:end], strings.Repeat(" ", pos))
	fmt.Println(e)
	return e
}

// GetObject finds an object by its qualified name, which looks like "location.latitude"
// as one example. Returns as interface{} to maintain generic handling
func getObject(objIn interface{}, qname string, level string) interface{} {
	if objIn == nil {
		fmt.Printf("Error: GetObject received nil schema from which to search for %s in %s\n", level, qname)
		os.Exit(1)
	}
	s := strings.SplitN(strings.TrimPrefix(level, "#/"), "/", 2)
	var leaf = len(s) == 1
	searchObj, found := objIn.(map[string]interface{})
	if !found {
		fmt.Printf("Error: object %s not map shaped at level %s\n", qname, level)
	}
	props, found := (searchObj["properties"]).(map[string]interface{})
	if !found {
		props, found = (searchObj["patternProperties"]).(map[string]interface{})
	}
	if found {
		searchObj = props
	}
	if o, found := searchObj[s[0]]; found {
		if leaf {
			return o
		}
		return getObject(o, qname, s[1])
	}
	return nil
}

// replaces all references recursively in the passed-in object (subschema) using the passed-in schema
func replaceReferences(schema map[string]interface{}, name string, obj interface{}) interface{} {
	oMap, isMap := obj.(map[string]interface{})
	switch {
	default:
		return obj
	case isMap:
		for k, v := range oMap {
			if k == "$ref" {
				r, found := lookup[v.(string)]
				if !found {
					fmt.Printf("** ERROR ** replaceReferences failed to lookup %s\n", v.(string))
					os.Exit(1)
				}
				if mapr, found := r.(map[string]interface{}); found {
					for kk, vv := range mapr {
						oMap[kk] = replaceReferences(schema, kk, vv)
					}
				}
				delete(oMap, "$ref")
			} else {
				oMap[k] = replaceReferences(schema, k, v)
			}
		}
		return oMap
	}
}

// Generates a file <munged elementName>.go to contain a string literal for the pretty version
// of the schema with all references resolved. In the same file, creates a sample JSON that
// can be used to show a complete structure of the object.
func generateGoSchemaFile(schema map[string]interface{}, config Config, imports string, regSchemas string) {
	var obj interface{}
	var schemas = make(map[string]interface{})
	var outString = "package main\n\n" + imports + "\n\n" + "var schemas = `\n\n"

	var filename = config.Schemas.GoSchemaFilename
	var apiFunctions = config.Schemas.API
	var elementNames = config.Schemas.Model

	schemas["API"] = make(map[string]interface{})
	schemas["Model"] = make(map[string]interface{})

	for i := range apiFunctions {
		functionSchemaName := "API/" + apiFunctions[i]
		obj = getObject(schema, functionSchemaName, functionSchemaName)
		if obj == nil {
			fmt.Printf("** WARN ** %s returned nil from getObject\n", functionSchemaName)
			return
		}
		schemas["API"].(map[string]interface{})[apiFunctions[i]] = obj
	}

	for i := range elementNames {
		elementName := "Model/" + elementNames[i]
		obj = getObject(schema, elementName, elementName)
		if obj == nil {
			fmt.Printf("** ERR ** %s returned nil from getObject\n", elementName)
			return
		}
		schemas["Model"].(map[string]interface{})[elementNames[i]] = obj
	}

	schemaOut, err := json.MarshalIndent(&schemas, "", "    ")
	if err != nil {
		fmt.Printf("** ERR ** cannot marshal schema file output for writing\n")
		return
	}
	outString += string(schemaOut) + "`\n\n" + regSchemas
	ioutil.WriteFile(filename, []byte(outString), 0644)
}

func sampleType(obj interface{}, elementName string) interface{} {
	o, found := obj.(map[string]interface{})
	if !found {
		return "SCHEMA ELEMENT " + elementName + " IS NOT MAP"
	}
	t, found := o["type"].(string)
	if !found {
		if elementName == "oneOf" {
			return o
		}
		return "NO TYPE PROPERTY"
	}
	switch t {
	default:
		fmt.Printf("** WARN ** Unknown type in sampleType %s\n", t)
	case "number":
		return 123.456
	case "integer":
		return 789
	case "string":
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
		desc, found := o["description"].(string)
		if found && len(desc) > 0 {
			return desc
		}
		return "carpe noctem"
	case "null":
		return nil
	case "boolean":
		return true
	case "array":
		var items, found = o["items"].(map[string]interface{})
		if !found {
			// fmt.Printf("** WARN ** Element %s is array with no items property\n", elementName)
			return "ARRAY WITH NO ITEMS PROPERTY"
		}
		return arrayFromSchema(items, elementName)
	case "object":
		{
			var props map[string]interface{}
			var found bool
			props, found = o["properties"].(map[string]interface{})
			if !found {
				props, found = (o["patternProperties"]).(map[string]interface{})
			}
			if !found {
				// fmt.Printf("** WARN ** %s is type object yet has no properties in SampleType\n", elementName)
				return "INVALID OBJECT - MISSING PROPERTIES"
			}
			objOut := make(map[string]interface{})
			for k, v := range props {
				//// fmt.Printf("Visiting key %s with value %s\n", k, v)
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
							//// fmt.Printf("SAMTYP outer OneOf: %d [%v]\n", k2, v2)
							vObj, found := v2.(map[string]interface{})
							if found {
								// inner loop should find one named object
								for k3, v3 := range vObj {
									tmp := make(map[string]interface{}, 1)
									//// fmt.Printf("SAMTYP inner OneOf: %s [%v]\n", k3, v3)
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

// Generate a sample array from a schema
func arrayFromSchema(schema map[string]interface{}, elementName string) interface{} {
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
func generateGoSampleFile(schema map[string]interface{}, config Config, imports string, regSamples string) {
	var obj interface{}
	var samples = make(map[string]interface{})
	var outString = "package main\n\n" + imports + "\n\n" + "var samples = `\n\n"

	var filename = config.Samples.GoSampleFilename
	var apiFunctions = config.Samples.API
	var modelNames = config.Samples.Model

	samples["API"] = interface{}(make(map[string]interface{}))
	samples["Model"] = interface{}(make(map[string]interface{}))

	for i := range apiFunctions {
		functionSchemaName := "API/" + apiFunctions[i]
		// use the schema subset
		obj = getObject(schema, functionSchemaName, functionSchemaName)
		if obj == nil {
			fmt.Printf("** WARN ** %s returned nil from getObject\n", functionSchemaName)
			return
		}
		samples["API"].(map[string]interface{})[apiFunctions[i]] = sampleType(obj, functionSchemaName)
	}
	for i := range modelNames {
		modelName := "Model/" + modelNames[i]
		// use the schema subset
		obj = getObject(schema, modelName, modelName)
		if obj == nil {
			fmt.Printf("** WARN ** %s returned nil from getObject\n", modelName)
			return
		}
		samples["Model"].(map[string]interface{})[modelNames[i]] = sampleType(obj, modelName)
	}
	samplesOut, err := json.MarshalIndent(&samples, "", "    ")
	if err != nil {
		fmt.Println("** ERR ** cannot marshal sample file output for writing")
		return
	}
	outString += string(samplesOut) + "`\n\n" + regSamples
	ioutil.WriteFile(filename, []byte(outString), 0644)
}

func loadModelTables(schema map[string]interface{}) {
	model, modelfound := schema["definitions"].(map[string]interface{})["Model"].(map[string]interface{})
	if !modelfound {
		fmt.Println("Warning: no Model section found in schema")
		os.Exit(1)
	} else {
		// all model entries as copies placed in lookup table
		for k, v := range model {
			lookup["#/definitions/Model/"+k] = DeepMergeMap(v.(map[string]interface{}), make(map[string]interface{}, 0))
		}
		// references in copies replaced
		for k := range model {
			lookup["#/definitions/Model/"+k] = replaceReferences(schema, k, lookup["#/definitions/Model/"+k])
		}
	}
	if *verbose {
		lookupfilename := "Model.lookup.table.json"
		fmt.Println("Writing model lookup table to: " + lookupfilename)
		_ = ioutil.WriteFile(lookupfilename, PrettyPrintBytes(lookup), 0744)
	}
}

func buildResolvedSchema(schema map[string]interface{}) map[string]interface{} {
	newSchema := make(map[string]interface{}, 0)
	newSchema["Model"] = make(map[string]interface{}, 0)
	var names []string
	for name, modelObj := range lookup {
		names = strings.SplitAfter(name, "#/definitions/Model/")
		if len(names) != 2 {
			fmt.Println("Error: cannot get last segment of name: " + name)
			os.Exit(1)
		}
		newSchema["Model"].(map[string]interface{})[names[1]] = modelObj.(map[string]interface{})
	}
	// use table to resolve API references
	newAPI := DeepMergeMap(schema["definitions"].(map[string]interface{})["API"].(map[string]interface{}), make(map[string]interface{}))
	newAPI = replaceReferences(schema, "API", newAPI).(map[string]interface{})
	newSchema["API"] = newAPI

	return newSchema
}

func getIncludedFile(path string, localpath string) string {
	var retstr = make([]byte, 0)
	var err error
	parts := strings.Split(path, "/#")
	includeschema := parts[0]
	level := parts[1]
	paths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, s := range paths {
		retstr, err = ioutil.ReadFile(s + "/src/" + includeschema)
		if err != nil {
			if os.IsNotExist(err) {
				// fmt.Println("****** include file not found on path: " + s + "/src/" + includeschema)
				continue
			}
			panic(errors.New("unknown error reading include file " + s + includeschema + err.Error()))
		} else {
			// fmt.Println("****** include file found on path: " + s + "/src/" + includeschema)
			break
		}
	}
	if retstr == nil || len(retstr) == 0 {
		fmt.Println("\nIncluded schema file not found on GOPATH: " + os.Getenv("GOPATH") + ": " + includeschema)
		fmt.Println("    -- please ensure that you have cloned or fetched the platform to your GOPATH")
		fmt.Println("    -- also, please ensure that you add /local-dev to your GOPATH using the command")
		fmt.Printf("       'export GOPATH=/opt/gopath:/local-dev' and then run 'go generate' again\n\n")
		os.Exit(1)
	}
	var ischema interface{}
	err = json.Unmarshal(retstr, &ischema)
	if err != nil {
		fmt.Println("*********** UNMARSHAL ERR **************\n", err)
		synerr := printSyntaxErrorIncludes(string(retstr), err)
		incfilename := localpath + "schema.with.failed.include.json"
		fmt.Println("Writing filed preprocessed schema to: " + incfilename)
		_ = ioutil.WriteFile(incfilename, retstr, 0744)
		panic(errors.New("unmarshal of schema with includes failed with err: " + synerr))
	}
	m, found := ischema.(map[string]interface{})
	if !found {
		panic(errors.New("included schema is not map shaped"))
	}
	o := getObject(m, "#/"+level, "#/"+level)
	if o == nil {
		panic(errors.New("Level " + level + " not found in schema " + path))
	}
	retstr, err = json.MarshalIndent(o, "", "   ")
	if err != nil {
		panic(errors.New("Level " + level + " in schema " + path + " failed to marshal: " + err.Error()))
	}
	// fmt.Println("\n\n RETURNING FROM INCLUDE PROCESSING\n\n" + string(retstr) + "\n\n")
	return string(retstr)
}

// Reads payloadschema.json api file
// encodes as a string literal in payloadschema.go
func main() {

	flag.Parse()

	if *verbose {
		fmt.Printf("genschema runs with config file %s\n", *configFile)
	}

	var regReadSamples = `
	var readAssetSamples iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(samples), nil
	}

	func init() {
		iot.AddRoute("readAssetSamples", "query", iot.SystemClass, readAssetSamples)
	}
	`
	var regReadSchemas = `
	var readAssetSchemas iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		iot.AddRoute("readAssetSchemas", "query", iot.SystemClass, readAssetSchemas)
	}
	`
	var imports = `
	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)`

	var api string
	var line = 1
	var lineOut = 1
	var offsets [5000]int

	filename, _ := filepath.Abs("./" + *configFile)
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(errors.New("error reading json file" + err.Error()))
	}
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		panic(errors.New("error unmarshaling json config" + err.Error()))
	}

	// ************** Stage 1
	// read the schema and preprocess for file includes at the top level
	filepre, err := os.Open(config.Schemas.SchemaFilename)
	if err != nil {
		fmt.Printf("** ERR ** [%s] opening input schema file at %s\n", err, config.Schemas.SchemaFilename)
		return
	}
	defer filepre.Close()
	reader := bufio.NewReader(filepre)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		// fmt.Println("MAIN SCHEMA: " + l)
		ts := strings.TrimSpace(l)
		if strings.HasPrefix(ts, "#") {
			fmt.Println("Line: ", line, " is a comment")
		} else if ts == "" {
			fmt.Println("Line: ", line, " is blank")
		} else if strings.HasPrefix(ts, "\"$ref\"") && strings.Index(ts, "\"#/") == -1 {
			ss := strings.Split(ts, "\"")
			p := ss[len(ss)-2]
			// fmt.Printf("line: %d includes: %s\n", line, p)
			refArr := getIncludedFile(p, "./")
			lines := strings.Split(refArr, "\n")
			// remove open and close brace as we are replacing the reference in place with the contents of the names object
			lines = lines[1 : len(lines)-1]
			for _, l2 := range lines {
				api += l2 + "\n"
				lineOut++
			}
			if len(ss) > 0 && ss[len(ss)-1] == "," {
				api += ","
			}
		} else {
			api += l + "\n"
			lineOut++
		}
		offsets[lineOut] = line
		line++
	}

	// ************** Stage 2
	// unmarshal the preprocessed schema
	var schema map[string]interface{}
	err = json.Unmarshal([]byte(api), &schema)
	if err != nil {
		fmt.Println("*********** UNMARSHAL ERR **************\n", err)
		printSyntaxErrorOffsets(api, &offsets, err)
		return
	}

	if *verbose {
		prefilename := strings.Split(filename, "/")[0] + "schema.with.includes.json"
		fmt.Println("Writing preprocessed schema to: " + prefilename)
		_ = ioutil.WriteFile(prefilename, PrettyPrintBytes(schema), 0744)
	}

	// ************** Stage 3
	// load the lookup tables with the data model, resolves all Model references
	loadModelTables(schema)

	// ************** Stage 4
	// build final schema by inserting API, resolving all references using lookup table, and
	// inserting the data model from the lookup table
	finalschema = buildResolvedSchema(schema)

	if *verbose {
		finalfilename := strings.Split(filename, "/")[0] + "schema.with.no.refs.json"
		fmt.Println("Writing final schema to: " + finalfilename)
		_ = ioutil.WriteFile(finalfilename, []byte(PrettyPrint(finalschema)), 0744)
	}

	// ************** Stage 5
	// generate the Go files that the contract needs -- for now, complete schema and
	// event schema and sample object

	generateGoSchemaFile(finalschema, config, imports, regReadSchemas)
	generateGoSampleFile(finalschema, config, imports, regReadSamples)

}
