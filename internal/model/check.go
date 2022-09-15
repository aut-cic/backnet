package model

type Check struct {
	ID        uint `gorm:"<-:false,autoIncrement"`
	Username  string
	Attribute string
	Op        string
	Value     string
}

func (Check) TableName() string {
	return "radcheck"
}
