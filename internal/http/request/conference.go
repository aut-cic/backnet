package request

type Conference struct {
	Name  string `form:"name"  json:"name"`
	Count int    `form:"count" json:"count"`
	Group string `form:"group" json:"group"`
}
