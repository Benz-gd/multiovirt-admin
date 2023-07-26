package controllers

import (
	"github.com/gin-gonic/gin"

)

const CtxUserIDKey = "userID"
func getCurrentUser(c *gin.Context)(userID int64,err error){
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		ResponseError(c,CodeUserNotLogin)
		return
	}
	userID,ok = uid.(int64)
	if !ok {
		ResponseError(c,CodeUserNotLogin)
		return
	}
	return userID,nil
}
