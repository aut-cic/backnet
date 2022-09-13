package request

type Conference struct {
	Name  string `json:"name" form:"name"`
	Count int    `json:"count" form:"count"`
	Group string `json:"group" form:"group"`
}
