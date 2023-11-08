package mysql

import (
	"errors"
	"go.uber.org/zap"
	"multiovirt-admin/models"
)

func ListHostGroup() (listhostgroup []*models.ListHostGroup, err error) {
	result := MysqlCMDB.Raw("select groupid,groupname,status,description,createdby,modifiedby,createdon,modifiedon from ListHostGroup").Scan(&listhostgroup)
	if result.Error != nil {
		zap.L().Error("ListHostGroup Error!")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("ListHostGroup Record Not Found!")
		return nil, errors.New("ListHostGroup Record Not Found!")
	}
	return listhostgroup, nil
}
