package model

type Check struct {
	ID        uint
	Username  string
	Attribute string
	Op        string
	Value     string
}

func (Check) TableName() string {
	return "radcheck"
}
