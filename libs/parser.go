package parser

//import _"encoding/json"
import ("fmt" 
		"errors"
		"os"
		"encoding/json"
		"io/ioutil"
		"gopkg.in/yaml.v2"
		"strings"
		"net/url"
		"net"
		"log"
		"flag"
		_"strings"
)


func ProcessArgs (args []string) (string, error) {

	if len(args) == 0 {
		return "", errors.New(ERROR_FILE_PATH_ARGUMENT_MISSING)
	} 

	if  len(args) == 1 {
		fmt.Println(" > Welcome to OpenAPI generator")
		return fmt.Sprintf("%s", args[0]), nil
	}

	if len(args) == 2 {
		fmt.Println(" > Welcome to Docker env variables extractor")

		var flagExtractor string;
		flag.StringVar(&flagExtractor, "--extractor", args[1], "extract the variables in your docker-compose environements")
	}

	return "", nil
	
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
	/*
	fmt.Println(collectionJson)

	for _, v := range(collectionJson.Item) {
		fmt.Println(strings.ToLower(v.Request.Method))
	}*/

	var swaggerInfo = SwaggerInfo {
		MetaDataApiVersion: INFO_API_VERSION_DEFAULT,
		MetaDataTitle: INFO_TILTE_DEFAULT,
		MetaDataDescription: INFO_DESCRIPTION_DEFAULT}

	var collectionYaml = CollectionYaml {
		SwaggerVersion: SWAGGER_VERSION,
		SwaggerInfo: swaggerInfo}

	
	chYaml := make(chan []byte, 1)
	go generatePathsContent(collectionYaml, collectionJson, chYaml)
	yamlUpdated := <-chYaml
	return string(yamlUpdated), nil

}


func generatePathsContent(collectionYaml CollectionYaml, collectionJson Collection, channel chan []byte) error {
	pathsContent := make(map[string]interface{})

	for _, v := range(collectionJson.Item) {
		parsedUrl := strings.ToLower(v.Request.Url.Raw)
		u, err := url.Parse(parsedUrl)

		if err != nil {
			log.Println("Error while parsing URL in collection postman")
		}

		host, port, _ := net.SplitHostPort(u.Host)
		// TODO add this to swagger host section
		log.Println(" > ", u.Path , " - host : ", host, " - port :", port )
		

		pathDetail := make(map[string]interface{})
		pathsContent[u.Path] = pathDetail

		methodDetail := make(map[string]interface{})
		pathDetail[strings.ToLower(v.Request.Method)] = methodDetail


	}

	collectionYaml.SwaggerPath = pathsContent


	yamlData, err := yaml.Marshal(&collectionYaml)

	if err != nil {
		return errors.New(ERROR_CONVERT_TO_YAML)
	}

	channel <- yamlData
	return nil
}

func extractVariablesDockerCompose() string {
	return "salut";
}
