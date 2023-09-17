package model


type (

	User struct {
		Uid          int `gorm:"primary_key;column:uid" json:"uid"`
		Username     string `gorm:"column:username" json:"username"`
		Password     string `gorm:"column:password" json:"password"`

		ExtInfo
	}

	ExtInfo struct {
		NewPassword string `gorm:"-" json:"new_password"`
	}
)

// TableName 设置表名
func (u User) TableName() string {
	return "user"
}

func NewUser(uid int, username, password string) *User {
	return &User{
		Uid: uid,
		Username: username,
		Password: password,
	}
}