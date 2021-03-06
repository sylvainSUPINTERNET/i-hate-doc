package parser

import "reflect"

const (
	ERROR_FILE_PATH_ARGUMENT_MISSING = "Missing file path in argument"
	ERROR_READ_JSON_COLLECTION_FILE = "Cannot parse the collection. File not found or wrong path given"
	ERROR_CONVERT_TO_YAML = "Error occured while trying to convert JSON postman collection to yaml format"
)


/**
	DEPRECATED
*/

type errorHandle struct {
	errorFilePathArgumentMissing string
}

func GetErrorHandler( errorKeyName string ) interface{} {
	errorHandle := errorHandle {errorFilePathArgumentMissing: "Missing file path in argument"}
	r := reflect.ValueOf(errorHandle)
	field := reflect.Indirect(r).FieldByName(errorKeyName)
	return field
}
