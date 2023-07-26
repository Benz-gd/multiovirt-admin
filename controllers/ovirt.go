package controllers

import (
	"example/fundemo01/web-app/logic"
	models "example/fundemo01/web-app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreateOvirtConf (c *gin.Context){
	//1、获取参数和参数校验
	ovirt := new(models.OvirtConf)
	if err := c.ShouldBindJSON(ovirt); err != nil {
		zap.L().Error("CreateOvirtCluster with invalid param", zap.Error(err))
		errT, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParam)
			return
		}else {
			ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errT.Translate(trans)))
			return
		}
		return
	}
	//2、检查参数是否在数据库存在
	if err := logic.CheckOvirtDBConf(ovirt);err != nil{
		ResponseError(c,CodeOvirtConfExist)
		return
	}
	//3、插入数据进入数据库
	rowAffected,err := logic.InsterOvirtDBConf(ovirt);if err != nil{
		ResponseError(c,CodeInsertDBErr)
		return
	}
	//3、返回响应

	ResponseSucWithMsgData(c,CodeRegisterSuccess,fmt.Sprintf("rowAffected %d",rowAffected))
}

func ListOvirtConf(c *gin.Context){
	listovirt,err := logic.ListOvirtConf()
	if err != nil{
		ResponseErrorWithMsg(c,CodeDataOperationErr,err)
		return
	}
	ResponseSucWithMsgData(c,CodeDataOperationSuccess,listovirt)
}

func GetOvirtConfDetail(c *gin.Context){
	//1、获取别名参数
	aliasname := c.Param("aliasname")
	if len(aliasname) == 0 {
		zap.L().Warn("OvirtConfDetail aliasname is null!")
		return
	}
	//2、根据别名查询所有字段
	ovirtconfdetail,err  := logic.GetOvirtConfDetail(aliasname)
	if err != nil{
		zap.L().Error("GetOvirtConfDetail ovirtconfdetail failed!",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	ResponseSucWithMsgData(c,CodeDataOperationSuccess,ovirtconfdetail)
}