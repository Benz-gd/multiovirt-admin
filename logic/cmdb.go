package logic

import (
	"multiovirt-admin/dao/mysql"
	"multiovirt-admin/models"
)

func ListHostGroup() (listhostgroup []*models.ListHostGroup, err error) {
	listhostgroup, err = mysql.ListHostGroup()
	if err != nil {
		//zap.L().Error("ListOvirtConf error!",zap.Error(err))
		return nil, err
	}
	return listhostgroup, nil
}
