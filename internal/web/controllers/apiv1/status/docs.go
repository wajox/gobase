package status

// ResponseDoc is a response declaration for documentatino pruposes
type ResponseDoc struct {
	Data struct {
		Attributes Response `json:"attributes"`
	} `json:"data"`
}
