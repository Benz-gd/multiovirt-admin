package middleware

import (
	"example/fundemo01/web-app/controllers"
	"example/fundemo01/web-app/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuthMiddleware() func(c *gin.Context){
	return func(c *gin.Context){
		//客户端携带的token有三种方式,1、放在请求头 2、放在请求体 3、放在URI
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == ""{
			controllers.ResponseError(c,controllers.CodeAuthNotExist)
			c.Abort()
			return
		}
		//按照空格分割
		parts := strings.SplitN(authHeader," ",2)
		if !(len(parts)== 2 && parts[0] == "Bearer"){
			controllers.ResponseError(c,controllers.CodeAuthFormatErr)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c,controllers.CodeAuthInvalidToken)
			c.Abort()
			return
		}
		c.Set(controllers.CtxUserIDKey,mc.UserID)
		c.Next()
	}
}
