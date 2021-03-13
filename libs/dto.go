package parser

/**
	Minimum struct to compose with collection from postman (JSON format)
*/

type Collection struct {
	Info Info `json:"info"`
	Item[] Item `json:"item"`
}


type Info struct {
	Name string `json:"name"`
}

type Item struct {
	Name string `json:"name"`
	Request Request `json:"request"`
	Response[] Response `json:"response"`
}

type Request struct {
	Method string `json:"method"`
	Header[] Header `json:"header"`
	Url Url `json:"url"`
}

// TODO
type Header struct {

}

type Url struct {
	Raw string `json:"raw"`
	Protocol string `json:"protocol"`
	Host[] string `json:"host"`
	Port string `json:"port"`
	Path[] string `json:"path"`
}

// TODO
type Response struct {

}