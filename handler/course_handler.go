package handler

import (
	"netdisk/model"
	"netdisk/service"

	"github.com/gin-gonic/gin"
)

var cs = service.NewCourseService()


func AddCourseHandler(c *gin.Context) {
	course := model.NewCourse(0, "")
	if err := c.ShouldBind(course); err != nil {
		return
	} 
	r := cs.AddCourse(course)
	c.JSON(200, r)
}

func JoinCourseHandler(c *gin.Context) {
	uc := model.NewUserCourse(0, 0)
	if err := c.ShouldBind(uc); err != nil {
		return
	}
	r := cs.JoinCourse(uc)
	c.JSON(200, r)
}