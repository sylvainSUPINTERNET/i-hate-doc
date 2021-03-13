package parser

//import _"encoding/json"
import ("fmt" 
		"errors"
		"os"
		"encoding/json"
		"io/ioutil"
		"gopkg.in/yaml.v2"
)


func ProcessArgs (args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New(ERROR_FILE_PATH_ARGUMENT_MISSING)
	} 
	return fmt.Sprintf("%s", args[0]), nil
}

func JsonCollectionParse (filePath string) (Collection, error) {
	var collection Collection

	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close() // reopen it later when required

	if err != nil {
		return collection, errors.New(ERROR_READ_JSON_COLLECTION_FILE)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &collection)
	return collection, nil
}

func JsonCollectionToYaml (collectionJson Collection) (string, error){
	var swaggerInfo = SwaggerInfo {
		MetaDataApiVersion: INFO_API_VERSION_DEFAULT,
		MetaDataTitle: INFO_TILTE_DEFAULT,
		MetaDataDescription: INFO_DESCRIPTION_DEFAULT}

	var collectionYaml = CollectionYaml {
		SwaggerVersion: SWAGGER_VERSION,
		SwaggerInfo: swaggerInfo,
		SwaggerPath: PATHS}


	yamlData, err := yaml.Marshal(&collectionYaml)
	
	if err != nil {
		return "", errors.New(ERROR_CONVERT_TO_YAML)
	}

	return string(yamlData), nil
}