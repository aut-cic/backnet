package model

type Package struct {
	Groupname string
}

func (Package) TableName() string {
	return "radpackages"
}
