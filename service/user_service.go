package service

import (
	"golang.org/x/crypto/bcrypt"
	"rongqin.cn/todo_task/model"
	"rongqin.cn/todo_task/pkg/utils"
	"rongqin.cn/todo_task/serialize"
)

type IUserSerivce interface {
	// 用户注册
	Register(formUserRegister *FormUserRegister) serialize.Response
	// 用户登录
	Login(formUserLogin *FormUserLogin) serialize.Response

	// 用户信息
	Get(id uint) serialize.Response
}

type UserService struct{}

// 用户登录
type FormUserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"passowrd" binding:"required"`
}

// 用户注册
type FormUserRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Company  string `form:"company" json:"company" biding:"required"`
}

type LoginUser struct {
	UserName string      `json:"username"`
	Compnay  string      `json:"company"`
	Token    utils.Token `json:"token"`
}

// 用户注册
func (userservice UserService) Register(formUserRegister *FormUserRegister) serialize.Response {

	user := model.User{
		Username: formUserRegister.Username,
		Password: formUserRegister.Password,
		Company:  formUserRegister.Company,
	}
	// 检查用户是否存在
	if user.Exist() {
		return serialize.BadReq("该用户已存在")
	}
	// 保存数据
	user.Insert()
	return serialize.OK()
}

// 用户登录
func (userservice UserService) Login(formUserLogin *FormUserLogin) serialize.Response {

	var user model.User
	user.GetUserByName(formUserLogin.Username)
	if user.ID == 0 {
		return serialize.BadReq("用户名或密码错误")
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formUserLogin.Password))
	if err != nil {
		return serialize.BadReq("用户名或密码错误")
	}
	// 颁发token
	var tokenService utils.TokenService
	token, err := tokenService.Encode(user.ID, user.Username)
	if err != nil {
		return serialize.Fail("token 颁发异常", err)
	}
	return serialize.OKData(&LoginUser{UserName: user.Username, Compnay: user.Company, Token: *token})
}

func (userService UserService) Get(uid int) serialize.Response {
	var user *model.User
	user.Get(uid)
	return serialize.OKData(user)
}
