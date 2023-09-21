package service

import (
	"fmt"
	"netdisk/dao"
	"netdisk/model"
	"netdisk/tool"
)

// UserService 用户功能
type UserService struct {
	userDao *dao.UserDao
	repoDao *dao.RepoDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: new(dao.UserDao),
		repoDao: new(dao.RepoDao),
	}
}

// UserRegister 用户注册
func (us *UserService) UserRegister(user *model.User) tool.Res {
	if code, b2 := tool.RedisGetVal(user.Email); b2 {
		if code == user.ACode {
			user.Username = user.Email
			// 口令加密
			user.Password = tool.Encrypt(user.Password)
			b := us.userDao.AddUser(user)
			if b {
				user.ACode = ""
				user = us.userDao.CheckUserInfo(user)
				
				// ! 分配私人仓库根目录
				rp := &model.PersonalRepository{
					Uid: user.Uid,
					ParentId: 0,
					IsDir: 1,
					SrcName: user.Email,
				}
				if us.repoDao.AddPersonalRepo(rp) {
					rp = us.repoDao.PersonalRepoInfo(rp)
					user.RootId = rp.PRid
					if us.userDao.UpdateRootId(user) {
						return tool.GetGoodResult(user)
					} else {
						return tool.GetBadResult("allocate personal repo failed")
					}
				} else {
					return tool.GetBadResult("allocate personal repo failed")
				}
			} else {
				return tool.GetBadResult("register failed")
			}
			
		} else {
			return tool.GetBadResult("code wrong")
		}
	} else {
		return tool.GetBadResult("register failed")
	}
}

// UserLogin 用户登录（密码）
func (us *UserService) UserLogin(user *model.User) tool.Res {
	// 口令加密
	user.Password = tool.Encrypt(user.Password)
	user.Username = user.Email
	validatedUser := us.userDao.ValidateUser(user)
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

// 验证码登录
func (us *UserService) UserLoginByEmail(u *model.User) tool.Res {
	if code, b := tool.RedisGetVal(u.Email); b {
		if u.ACode == code {
			u.Username = u.Email
			validatedUser := us.userDao.CheckUserInfo(u)
			validatedUser.Password = "*****************"
			// 设置 token
			if token, err := tool.GetToken(fmt.Sprintf("%d%s", validatedUser.Uid, validatedUser.Username)); err != nil {
				return tool.GetBadResult("get Token failed")
			} else {
				return tool.GetGoodResult(*validatedUser, token)
			}
		} else {
			return tool.GetBadResult("code wrong, login failed")
		}
	} else {
		return tool.GetBadResult("login failed")
	}
}

// 用户注销
func (us *UserService) DeleteUser(user *model.User) tool.Res {
	if b := us.userDao.DelUser(user); b {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("failed")
	}
}

// 更改密码
func (us *UserService) UpdatePassword(user *model.User) tool.Res {
	if code, b := tool.RedisGetVal(user.Email); b {
		if code == user.ACode {
			user.Password = tool.Encrypt(user.NewPassword)
			if us.userDao.UpdatePassword(user)	{
				return tool.GetGoodResult(nil)
			} else {
				return tool.GetBadResult("err")
			}
		} else {
			return tool.GetBadResult("code wrong")
		}
	} else {
		return tool.GetBadResult("update failed")
	}
}

func (us *UserService)SendAuthenticationCodeEmail(user *model.User, branch int) tool.Res {
	switch branch {
	case 0: // !注册
		if us.userDao.CheckEmail(user) {
			return tool.GetBadResult("user exists")
		} else {
			acode := tool.CreateCaptcha(6)
			if tool.RedisSetkeyVal(user.Email, acode) {
				if err := tool.RegisterByEmail(user.Email, acode); err != nil {
					return tool.GetBadResult("send email failed")
				} else {
					return tool.GetGoodResult("发送成功")
				}			
			} else {
				return tool.GetBadResult("send email failed")
			}
		}
	case 1: // !登录
		if us.userDao.CheckEmail(user) {
			acode := tool.CreateCaptcha(6)
			if tool.RedisSetkeyVal(user.Email, acode) {
				if err := tool.LoginByEmail(user.Email, acode); err != nil {
					return tool.GetBadResult("send email failed")
				} else {
					return tool.GetGoodResult("发送成功")
				}			
			} else {
				return tool.GetBadResult("send email failed")
			}
		} else {
			return tool.GetBadResult("user does not exist")
		}
	case 2: // !找回密码，更改密码
		if us.userDao.CheckEmail(user) {
			acode := tool.CreateCaptcha(6)
			if tool.RedisSetkeyVal(user.Email, acode) {
				if err := tool.ResetPasswordEmail(user.Email, acode); err != nil {
					return tool.GetBadResult("send email failed")
				} else {
					return tool.GetGoodResult("发送成功")
				}			
			} else {
				return tool.GetBadResult("send email failed")
			}
		} else {
			return tool.GetBadResult("user does not exist")
		}
	default:
		return tool.GetBadResult("unkown")
	}
}