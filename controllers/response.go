package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int           `json:"code"`
	Msg interface{}    `json:"msg"`
	Data interface{}   `json:"data"`
}

func ResponseError(c *gin.Context,code int){
	rd := &ResponseData{
		Code: code,
		Msg: GetErrMsg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}


func ResponseSuccess(c *gin.Context,code int){
	rd := &ResponseData{
		Code: code,
		Msg: GetSucMsg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}


func ResponseErrorWithMsg(c *gin.Context,code int,msg interface{}){
	rd := &ResponseData{
		Code: code,
		Msg: msg,
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}


func ResponseSucWithMsgData(c *gin.Context,code int,data interface{}){
	rd := &ResponseData{
		Code: code,
		Msg: GetSucMsg(code),
		Data: data,
	}
	c.JSON(http.StatusOK,rd)
}
