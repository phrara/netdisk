package service

import (
	"netdisk/dao"
	"netdisk/model"
	"netdisk/tool"
	"strings"
)

type CourseService struct {
	courseDao *dao.CourseDao
}


func NewCourseService() *CourseService {
	cs := new(CourseService)
	cs.courseDao = new(dao.CourseDao)
	return cs
}


// 添加课程
func (cs *CourseService) AddCourse(c *model.Course) tool.Res {
	c.CourseName = strings.Trim(c.CourseName, " ")
	if tool.WordsInspect(c.CourseName) {

		if cs.courseDao.AddCourse(c) {
			return tool.GetGoodResult(*c)
		} else {
			return tool.GetBadResult("add course failed")
		}
	} else {
		return tool.GetBadResult("illeagal words")
	}
}

// 加入课程
func (cs *CourseService) JoinCourse(uc *model.UserCourse) tool.Res {
	if cs.courseDao.JoinCourse(uc) {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("join course failed")
	}
}