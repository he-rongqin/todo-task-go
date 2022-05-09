package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"rongqin.cn/todo_task/pkg/utils"
	"rongqin.cn/todo_task/serialize"
)

const (
	HEADER_AUTHORIZATION_KEY = "Authorization"
	TOKEN_PREFIX             = "Bearer "
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader(HEADER_AUTHORIZATION_KEY)
		if token == "" {
			ctx.JSON(serialize.NOT_AUTHENTICATION, serialize.Response{
				Status: serialize.NOT_AUTHENTICATION,
				Msg:    "无权限访问该资源，请登录之后再试",
			})
			// 拒绝请求
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(token, TOKEN_PREFIX) {
			ctx.JSON(serialize.BADREQUEST, serialize.Response{
				Status: serialize.BADREQUEST,
				Msg:    "token 格式错误",
			})
			// 拒绝请求
			ctx.Abort()
			return
		}
		// 解析token
		var tokenService utils.TokenService
		claims, err := tokenService.Decode(strings.Split(token, " ")[1])
		if err != nil {
			ctx.JSON(serialize.BADREQUEST, serialize.Response{
				Status: serialize.BADREQUEST,
				Msg:    "token 解析错误",
				ErrMsg: err.Error(),
			})
			// 拒绝请求
			ctx.Abort()
			return
		}
		// 令牌过期
		if time.Now().Unix() > claims.ExpiresAt {
			ctx.JSON(serialize.NOT_AUTHENTICATION, serialize.Response{
				Status: serialize.NOT_AUTHENTICATION,
				Msg:    "令牌已过期，请重新登录",
			})
			// 拒绝请求
			ctx.Abort()
			return
		}
		// todo 刷新token

		ctx.Next()

	}
}
