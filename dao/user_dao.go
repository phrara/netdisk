package dao

import (
	"netdisk/model"
)

type UserDao struct {
}

// CheckUserInfo 查询用户信息
func (ud *UserDao) CheckUserInfo(user *model.User) *model.User {
	u := model.NewUser(0, "", "")
	DBMgr.Where("username = ? or email = ?", user.Username, user.Email).First(u)
	return u
}

func (ud *UserDao) CheckEmail(user *model.User) bool {
	u := model.NewUser(0, "", "")
	DBMgr.Where("uid = ? and email = ?", user.Uid, user.Email).First(u)
	if u.Username != "" {
		return true
	} else {
		return false
	}
}

// ValidateUser 登录验证
func (ud *UserDao) ValidateUser(user *model.User) *model.User {
	u := model.NewUser(0, "", "")
	DBMgr.Where("username = ? and password = ?", user.Username, user.Password).First(u)
	return u
}

// AddUser 添加新用户
func (ud *UserDao) AddUser(user *model.User) bool {
	res := DBMgr.Create(user)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// 删除用户
func (ud *UserDao) DelUser(u *model.User) bool {
	
	DBMgr.Delete(u)
	if ud.CheckUserInfo(u).Username == "" {
		return true
	} else {
		return false
	}
}

// 更改密码
func (ud *UserDao) UpdatePassword(user *model.User) bool {
	res := DBMgr.Model(user).Where("uid = ?",user.Uid).Update("password",user.Password)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}


// 查询所有用户
func (ud *UserDao) GetAllUsers(num int) []model.User {
	ulist := make([]model.User, num)
	DBMgr.Order("registerTime desc").Limit(num).Find(&ulist)
	return ulist
}