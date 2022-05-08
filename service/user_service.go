package service

import (
	"rongqin.cn/todo_task/model"
	response "rongqin.cn/todo_task/serialize"
)

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

// 用户注册
func (formUserRegister *FormUserRegister) Register() response.Response {

	user := model.User{
		Username: formUserRegister.Username,
		Password: formUserRegister.Password,
		Company:  formUserRegister.Company,
	}
	// 检查用户是否存在
	if user.Exist() {
		return response.BadReq("该用户已存在")
	}
	// 保存数据
	user.Insert()
	return response.OK()
}
