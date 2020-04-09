package status

type ResponseDoc struct {
	Data struct {
		Attributes Response `json:"attributes"`
	} `json:"data"`
}
