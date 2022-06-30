package rbac

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liyinda/google-authenticator/utils/errno"
	myJwt "github.com/liyinda/google-authenticator/utils/jwt"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errno.SUCCESS
		//token := c.Query("token")
		//vue-admin 需要通过HTTP HEADER中取到X-Token代替之前的从url参数
		xtoken := c.Request.Header["X-Token"]
		token := xtoken[0]

		if token == "" {
			code = errno.INVALID_PARAMS
		} else {
			_, err := myJwt.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errno.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = errno.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != errno.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errno.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取前端传回的token(传递方式不同，获取的位置也不同，根据实际情况选择)
		//authHeader := ctx.Request.Header.Get(myJwt.AcceptTokenKey)
		//authHeader := ctx.Query("token")
		//vue-admin 需要通过HTTP HEADER中取到X-Token代替之前的从url参数
		xtoken := ctx.Request.Header["X-Token"]
		authHeader := xtoken[0]
		// 无token直接返回错误
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "未登录或非法访问",
			})
			// 校验失败终止后续操作
			ctx.Abort()
			return
		}
		// 解析token
		claims, err := myJwt.ParseToken(authHeader)
		// 错误处理
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": errno.ErrToken,
				"msg":  err.Error(),
				"data": nil,
			})
			ctx.Abort()
			return
		}
		// 将claim加入上下文，便于后续使用
		ctx.Set("userClaim", claims)
		ctx.Next()
	}
}

//用户会话保持
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}
