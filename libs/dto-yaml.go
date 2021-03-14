package parser

/**
	DEPRECATED
**/


type CollectionYaml struct {
	SwaggerVersion string `yaml:"swagger"`
	SwaggerInfo SwaggerInfo `yaml:"info"`
	SwaggerPath map[string]interface{} `yaml:"paths"`
}


type SwaggerInfo struct {
	MetaDataApiVersion string `yaml:"version"`
	MetaDataTitle string `yaml:"title"`
	MetaDataDescription string `yaml:"description"`
}
