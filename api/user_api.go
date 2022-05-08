package api

import (
	"github.com/gin-gonic/gin"
	response "rongqin.cn/todo_task/serialize"
	"rongqin.cn/todo_task/service"
)

// 用户登录
func UserLogin(c *gin.Context) {

}

// 用户注册
func UserRegister(c *gin.Context) {

	var userRegister service.FormUserRegister
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(response.BADREQUEST, response.BadReq(err.Error()))
		return
	}
	res := userRegister.Register()
	if res.Status != 200 {
		c.JSON(response.BADREQUEST, res)
		return
	}
	c.JSON(response.SUCCESS, res)

}
