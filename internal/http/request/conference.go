package request

type Conference struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Group string `json:"group"`
}
