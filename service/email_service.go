package service

import (
	"netdisk/dao"
	"netdisk/model"
	"netdisk/tool"
)

type EmailService struct {
	userdao *dao.UserDao
}

func NewEmailService() *EmailService {
	return &EmailService{
		userdao: new(dao.UserDao),
	}
}

func (es *EmailService) AuthenticationCodeEmail(user *model.User) tool.Res {
	if es.userdao.CheckEmail(user) {
		// TODO
	} else {
		return tool.GetBadResult("the email does not match")
	}
	return nil
}