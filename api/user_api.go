package api

import (
	"github.com/gin-gonic/gin"
	response "rongqin.cn/todo_task/serialize"
	"rongqin.cn/todo_task/service"
)

var userService service.UserService

// 用户登录
func UserLoginEndpoint(c *gin.Context) {

	var userLogin service.FormUserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(response.BADREQUEST, response.BadReq(err.Error()))
		return
	}
	res := userService.Login(&userLogin)
	if res.Status != 200 {
		c.JSON(response.BADREQUEST, res)
		return
	}
	c.JSON(response.SUCCESS, res)
}

// 用户注册
func UserRegisterEndpoint(c *gin.Context) {

	var userRegister service.FormUserRegister
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(response.BADREQUEST, response.BadReq(err.Error()))
		return
	}
	res := userService.Register(&userRegister)
	if res.Status != 200 {
		c.JSON(response.BADREQUEST, res)
		return
	}
	c.JSON(response.SUCCESS, res)

}

func GetUserInfoEndpoint(c *gin.Context) {

	u := c.GetInt("uid")
	// uid, _ := strconv.ParseInt(u, 6, 12)
	res := userService.Get(u)
	if res.Status != 200 {
		c.JSON(response.BADREQUEST, res)
		return
	}
	c.JSON(response.SUCCESS, res)
}
