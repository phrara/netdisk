package model


type UserCourse struct {
	Uid int `gorm:"column:uid" json:"uid"`
	Cid int `gorm:"column:cid" json:"cid"`
}

func (u UserCourse) TableName() string {
	return "user_course"
}

func NewUserCourse(uid, cid int) *UserCourse {
	return &UserCourse{
		Uid: uid,
		Cid: cid,
	}
}