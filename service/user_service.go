package service

import (
	"fmt"
	"netdisk/dao"
	"netdisk/model"
	"netdisk/tool"
	"strings"
)

// UserService 用户功能
type UserService struct {
	userdao *dao.UserDao
}

func NewUserService() *UserService {
	us := new(UserService)
	us.userdao = new(dao.UserDao)
	return us
}

// UserRegister 用户注册
func (us *UserService) UserRegister(user *model.User) tool.Res {
	user.Username = strings.Trim(user.Username, " ")
	if tool.WordsInspect(user.Username) {
		// 口令加密
		user.Password = tool.Encrypt(user.Password)
		b := us.userdao.AddUser(user)
		if b {
			return tool.GetGoodResult(*user)
		} else {
			return tool.GetBadResult("register failed")
		}
	} else {
		return tool.GetBadResult("illegal words")
	}

}

// UserLogin 用户登录
func (us *UserService) UserLogin(user *model.User) tool.Res {
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	validatedUser := us.userdao.ValidateUser(user)
	if validatedUser.Username != "" {
		validatedUser.Password = "*************"
		// 设置 token
		token, err := tool.GetToken(fmt.Sprintf("%d%s", validatedUser.Uid, validatedUser.Username))
		if err != nil {
			return tool.GetBadResult("get Token failed")
		}
		return tool.GetGoodResult(*validatedUser, token)
	} else {
		return tool.GetBadResult("login failed")
	}
}

// 用户注销
func (us *UserService) DeleteUser(user *model.User) tool.Res {
	if b := us.userdao.DelUser(user); b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}

// 更改密码
func (us *UserService) UpdatePassword(u *model.User, newPassword string) tool.Res {

	u.Password = tool.Encrypt(u.Password)
	if us.userdao.ValidateUser(u).Username == "" {
		return tool.GetBadResult("pwd err")
	}

	u.Password = tool.Encrypt(newPassword)

	b := us.userdao.UpdatePassword(u)	
	if b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("err")
	}
}