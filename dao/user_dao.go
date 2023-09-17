package dao

import (
	"netdisk/model"
)

type UserDao struct {
}

// GetUserInfo 查询用户信息
func (*UserDao) GetUserInfo(username string) *model.User {
	user := model.NewUser(0, "", "")
	DBMgr.Where("username = ?", username).First(user)
	return user
}

// ValidateUser 登录验证
func (*UserDao) ValidateUser(user *model.User) *model.User {
	u := model.NewUser(0, "", "")
	DBMgr.Where("username = ? and password = ?", user.Username, user.Password).First(u)
	return u
}

// AddUser 添加新用户
func (ud *UserDao) AddUser(user *model.User) bool {
	if ud.GetUserInfo(user.Username).Username == "" {
		DBMgr.Create(user)
		return true
	} else {
		return false
	}
}

// 删除用户
func (ud *UserDao) DelUser(u *model.User) bool {
	
	DBMgr.Delete(u)
	if ud.GetUserInfo(u.Username).Username == "" {
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