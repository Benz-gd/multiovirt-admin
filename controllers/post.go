package controllers

import (
	"example/fundemo01/web-app/logic"
	"example/fundemo01/web-app/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePost(c *gin.Context){
	// 1、获取表单的内容
	postdetail := new(models.ParamPost)
	if err := c.ShouldBindJSON(postdetail);err != nil{

	}

	//2、创建帖子


	//3、返回响应


	data,err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GegCommunityList() failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSucWithMsgData(c,CodeResponseSuccess,data)
}