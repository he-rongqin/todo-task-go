package router

import (
	"github.com/gin-gonic/gin"
	"rongqin.cn/todo_task/api"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
	}

	return router

}
