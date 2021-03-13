package parser

//import _"encoding/json"
import ("fmt" 
		"errors"
		"os"
		"encoding/json"
		"io/ioutil"
		_"gopkg.in/yaml.v2"
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
	return "", nil
}