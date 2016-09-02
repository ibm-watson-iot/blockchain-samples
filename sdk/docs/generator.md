# The Schema Generator Script

[`scripts/generate_go_schema.go`](scripts/generate_go_schema.go "generate selected portions of the contract schema in go")  
[`scripts/generate.json`](scripts/generate.json "configure the contract schema generator")

## Introduction

This module contains a Go program that acts as a script, executed when the command `go generate` is run while in the main contract folder. The module resides in the scripts folder because it is also main program (as in `package main`) and so would conflict with the contract's main package were they colocated.

The module performs three functions:

- validates the schema using the `json.Unmarshal` function and reports the first detailed error it finds at the command line 
- generates the file `schemas.go` in package main
- generates the file `samples.go` in package main

See [`schema.md`](schema.md "describes the payload schema and how it is used") for more information on commands, errors and outputs.

See [Understanding JSON Schema](http://spacetelescope.github.io/understanding-json-schema/UnderstandingJSONSchema.pdf "short reference document on JSON Schema 4") for a reasonably short, yet thorough reference on JSON Schema 4.

## Mainline Processing

### func main

The initial section concerns itself with loading its configuration from the json configuration file. 

>*Note that this could have been a yaml file for better readability, but we were unable to find a yaml project that did not use unnecessarily restrictive licenses, and json is human-readable enough for our purposes.* 

The next section removes blank lines and comments from the schema and generates an offset table to align line numbers in the processed schema with the correct line in the original `payloadSchema.json` file. This might be useful to some because comments are not allowed in JSON files, yet commenting the schema may be of value. We dropped comments from our schema for better compatibility with third party tools like Swagger, although we no longer find it necessary to use Swagger as the script is perfectly adequate during development. 

Next, the main function validates the schema by unmarshaling it. If an error is returned, it is passed to the function `printSyntaxError` along with the compact schema as a string and the offset table. The main function exits after printing the schema syntax error.

Finally, a new version of the schema is generated from the unmarshaled schema by replacing all references with their content, basically expanding the schema to eliminate references.

This new fully expanded schema is passed into the functions `generateGoSchemaFile` and `generateGoSampleFile` along with the config that was previously read and unmarshaled.

---

## Schemas generation

### func getObject

Retrieves an object based on a path. Assumes that the paths are relative to `#/definitions/` and so that does not need to appear. References inside the schema will generally have the full path. Used by replace references to perform the recursive schema expansion.

### func replaceReferences

Crawls the entire schema, replacing any `$ref` by the referenced content. 

### func referencesExist

Returns true if a reference exists anywhere in the passed in subschema. Deprecated and no longer used, but left in for potential future uses.

### func generateGoSchemaFile

Accepts as input the fully expanded schema and the configuration that specifies the APIs and Object Model Elements to include in the go file. 

The `schemas.go` file begins with the package and a string definition:

``` go

package main

var schemas = `
{
    "API": {
        "createAsset": {

```

The file is automatically visible to the contract by its membership in the main package, and the string defined therein can be returned exactly as defined to an application calling the API `getAssetSchemas`. The returned string is a JSON object with two contained objects, `API` and `objectModelSchemas`. (And yes, the name should have had more of a parallel feel for API, something like ObjectModel.) 

The API section contains the APIs named in 
[`generate.json`](../scripts/generate.json "configuration for schemas and samples processing"), co-located in the scripts folder beside this file:

``` json

    "API": [
      "init",
      "createAsset",
      "updateAsset",
      "deleteAsset",
      "deletePropertiesFromAsset",
      "deleteAllAssets",
      "readAsset",
      "readAllAssets",
      "readAssetHistory",
      "readRecentStates",
      "setLoggingLevel",
      "setCreateOnUpdate"
    ],

```

We generally include all CRUD APIs (create, read, update, delete) along with the init and the `set` APIs, which are used to control execution features of the contract such as logging and *update redirect to create on asset not found*.

Including the entire API has some advantages when dynamic mapping and form generation are considered. Plugins exist that allow generic GUIs to be created with minimal effort. One such is used by the simple and generic UIs in this project.

> Note that any forms that are used to create events to the contract should adhere strictly to the *partial state as event* pattern, which means that empty and zero values are **NOT** sent in messages to the contract. The contract distinguishes property existence by presence and absence in the message, **NOT** by normal or zero values. For example, the zero value for a JSON number is, obviously, 0. But that is also a perfectly valid temperature reading, and thus could not be used to define absence inside the contract under any circumstances. *Send only properties that contain real data and can replace previous values in the asset's state.*

Inputs and outputs are a part of the API definitions, so theoretically it is unnecessary to process the object model independently for applications that just want to send events and read state. 

But there is value to seeing the contract data model as a separate concern and so the script generates `objectModelSchemas` from the config entry `goSchemaElements` (another name that needs cleaning up in the future).

---

## Samples Generation

### func sampleType

This is the primary function in the simple sample generator. It is passed a schema object (remember that all elements are defined by schema objects with properties like `type`, `default`, `description` and so on) and an elementName. It attempts to find a type field and prints an error if type is missing. Note that you must fix this error before continuing as all JSON Schema 4 objects must have a type.

>*Why did the the initial error checking not spot this issue? The error checking performed in the mainline processing handled missing schema file and syntax errors only. Any valid JSON would pass that test. But only JSON Schema 4 compatible JSON will pass this test.*

If the type field is found, then a `switch` is performed on its value. Valid cases are:

- `number` : a JSON number is a float64 or equivalent and the generated sample value is always 123.456 
- `integer` : a JSON integer is an int or equivalent and the generated sample value is always 789
- `string` : a JSON string is a string and the generated sample value is one of:
  - an RFC3339nanos formatted timestamp if the name of the field is `timestamp` (case independent)
  - the contents of the `example` field if found
  - the contents of the `default` field if found
  - if an `enum` field is found
    - the *second* entry in an enum of length > 1
    - else the only entry  
      > *Why the second entry? Because the first entry is often some form of nil, and the last might be something indicating the length (as in <enumname>_size), so choosing the second entry if possible is the safest*
  - the `description` field if found
  - `"carpe noctem" if nothing else is available`
- `null` : a JSON null maps to a sample value of `nil`
-  `boolean` : a JSON boolean maps to a sample value of `true`
- `array` : gets a sample value from the function `arrayFromSchema`
- `object` : a nested JSON object gets sample values recursively

### func arrayFromSchema

Returns an array `enum` if present, else returns an array with a sample of the element inside the `items` object.

### generateGoSampleFile

Accepts as input the fully expanded schema and the configuration that specifies the APIs and Object Model Elements to include in the go file. 

The `samples.go` file begins with the package and a string definition:

``` go

package main

var samples = `
{
    "contractState": {
        "activeAssets": [

```

The returned string is a JSON object with samples for each of the specified objects:

``` json

    "goSampleElements": [
      "initEvent",
      "event",
      "state",
      "contractState"
    ]

```

Samples can be useful to preload objects in a (for example) JavaScript client. Scenarios will exist where calling `getAssetSamples` for the preconfigured objects is preferred to calling `getAssetSchemas` and processing the raw JSON Schema.

## Conclusion

There are a few other bits of code in the script that are not in use currently. They exist for easy testing of the script (by uncommenting certain lines in main) and so are left for the convenience of contract developers.