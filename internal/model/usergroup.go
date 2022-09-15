package model

type UserGroup struct {
	Username  string
	Groupname string
	ID        uint `gorm:"<-:false,autoIncrement"`
}

func (UserGroup) TableName() string {
	return "radusergroup"
}
