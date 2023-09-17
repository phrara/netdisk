package model

type Course struct {
	Cid int `gorm:"primary_key;column:cid" json:"cid"`
	CourseName string `gorm:"column:course_name" json:"course_name"`
}

func (c Course) TableName() string {
	return "course"
}	


func NewCourse(cid int, coursename string) *Course {
	return &Course{
		Cid: cid,
		CourseName: coursename,
	}
}