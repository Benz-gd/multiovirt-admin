package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"multiovirt-admin/logic"
	"multiovirt-admin/models"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	//1 获取参数和参数校验
	paramsignup := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(paramsignup); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是否validator.ValidationErrors类型
		errT, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errT.Translate(trans)))
			return
		}
		return
	}
	//2、业务处理
	if err := logic.SignUp(paramsignup); err != nil {
		ResponseError(c, CodeRegisterFail)
		return
	}
	//3、返回响应
	ResponseSuccess(c, CodeRegisterSuccess)
}

func LoginHandler(c *gin.Context) {
	//1、获取请求参数及参数处理
	paramlogin := new(models.ParamLogin)
	if err := c.ShouldBindJSON(paramlogin); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是否validator.ValidationErrors类型
		errT, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errT.Translate(trans)),
			})
			return
		}
		return
	}
	//2、业务逻辑处理
	if atoken, _, err := logic.Login(paramlogin); err != nil {
		zap.L().Error("Login with invalid param", zap.String("username", paramlogin.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户或密码错误!",
		})
		return
	} else {
		//3、返回响应

		c.JSON(http.StatusOK, gin.H{
			"msg":  "登录成功!",
			"data": atoken,
		})
	}
}
