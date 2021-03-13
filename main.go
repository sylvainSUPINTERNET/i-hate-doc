package main

import (
	"fmt"
	"os"
	"github.com/sylvainSUPINTERNET/openapi-generator/libs"
)

func main() {	
	fmt.Println(" > Welcome to OpenAPI generator")
	filePath, err := parser.ProcessArgs(os.Args[1:])

	if err != nil {
		fmt.Println(fmt.Sprintf(" ! * %s", err))
		os.Exit(1)
	}

	collectionJson, errParseJsonToCollection := parser.JsonCollectionParse(filePath)

	if errParseJsonToCollection != nil {
		fmt.Println(fmt.Sprintf(" ! * %s", errParseJsonToCollection))
		os.Exit(1)
	}

	parser.JsonCollectionToYaml(collectionJson)
}