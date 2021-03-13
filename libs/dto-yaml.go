package parser


type CollectionYaml struct {
	SwaggerVersion string `yaml:"swagger"`
	SwaggerInfo SwaggerInfo `yaml:"info"`
	SwaggerPath string `yaml:"path"`
}


type SwaggerInfo struct {
	MetaDataApiVersion string `yaml:"version"`
	MetaDataTitle string `yaml:"title"`
	MetaDataDescription string `yaml:"description"`
}