package dao

import "netdisk/model"

type CourseDao struct {
}

// 查询课程
func (cd *CourseDao) GetCourseInfo(courseName string) *model.Course {
	course := model.NewCourse(0, "")
	DBMgr.Where("course_name = ?", courseName).First(course)
	return course
}

// 添加课程
func (cd *CourseDao) AddCourse(c *model.Course) bool {
	if cd.GetCourseInfo(c.CourseName).CourseName == "" {
		DBMgr.Create(c)
		return true
	} else {
		return false
	}
}


// 加入课程
func (cd *CourseDao) JoinCourse(c *model.UserCourse) bool {
	uc := model.NewUserCourse(0, 0)
	DBMgr.Where("cid = ? and uid = ?", c.Cid, c.Uid).First(uc)
	if uc.Cid == 0 && uc.Uid == 0 {
		DBMgr.Create(c)
		return true
	} else {
		return false
	}
}