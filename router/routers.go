package router

import (
	"github.com/gin-gonic/gin"
	"rongqin.cn/todo_task/api"
	"rongqin.cn/todo_task/middleware"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegisterEndpoint)
		v1.POST("user/login", api.UserLoginEndpoint)
	}

	authorized := router.Group("/api/v1")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/user/info", api.GetUserInfoEndpoint)
	}
	return router

}
