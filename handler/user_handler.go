package handler

import (
	"netdisk/model"
	"netdisk/service"

	"github.com/gin-gonic/gin"
)

var us = service.NewUserService()


func UserRegisterHandler(c *gin.Context) {
	user := model.NewUser(0, "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	res := us.UserRegister(user)
	c.JSON(200, res)
}


func UserLoginHandler(c *gin.Context) {
	user := model.NewUser(0, "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	b := us.UserLogin(user)
	c.JSON(200, b)
}

func UserLoginByEmailHandler(c *gin.Context) {
	user := model.NewUser(0, "", "")
	err := c.ShouldBind(user)
	if err != nil {
		return
	}
	b := us.UserLoginByEmail(user)
	c.JSON(200, b)
}


// 用户注销
func DelUserHandler(c *gin.Context){
	u := new(model.User)
	err := c.ShouldBind(u)
	if err != nil {
		return
	}
	b := us.DeleteUser(u)
	c.JSON(200, b)
}

// 修改密码
func UpdatePasswordHandler(c *gin.Context){
	u := new(model.User)
	err := c.ShouldBind(u)
	if err != nil {
		return
	}
	b := us.UpdatePassword(u)
	c.JSON(200, b)
}


func RegisterSendCodeEmail(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBind(u); err != nil {
		return
	}
	r := us.SendAuthenticationCodeEmail(u, 0)
	c.JSON(200, r)
}
func LoginSendCodeEmail(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBind(u); err != nil {
		return
	}
	r := us.SendAuthenticationCodeEmail(u, 1)
	c.JSON(200, r)
}
func UpdatePasswordSendCodeEmail(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBind(u); err != nil {
		return
	}
	r := us.SendAuthenticationCodeEmail(u, 2)
	c.JSON(200, r)
}