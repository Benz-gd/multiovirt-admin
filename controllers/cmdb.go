package controllers

import (
	"github.com/gin-gonic/gin"
	"multiovirt-admin/logic"
)

func ListHostGroup(c *gin.Context) {
	listcmdbgroups, err := logic.ListHostGroup()
	if err != nil {
		ResponseErrorWithMsg(c, CodeDataOperationErr, err)
		return
	}
	ResponseSucWithMsgData(c, CodeDataOperationSuccess, listcmdbgroups)
}
