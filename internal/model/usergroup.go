package model

type UserGroup struct {
	Username  string
	Groupname string
	ID        uint
}

func (UserGroup) TableName() string {
	return "radusergroup"
}
