package controllers

import (
	"example/fundemo01/web-app/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityList(c *gin.Context){
	// 查询到所有的数据 (community_id,community_name)以列表的形式放回
	data,err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GegCommunityList() failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSucWithMsgData(c,CodeResponseSuccess,data)
}


func CommunityDetail(c *gin.Context){
	// 1、获取参数
	communityID,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		zap.L().Error("CommunityDetail cannot change int!",zap.Error(err))
		return
	}

	// 2、根据ID查询到所有字段
	data,err := logic.GetCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.GegCommunityList() failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	ResponseSucWithMsgData(c,CodeResponseSuccess,data)
}



