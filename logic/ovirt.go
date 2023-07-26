package logic

import (
	"example/fundemo01/web-app/dao/mysql"
	"example/fundemo01/web-app/models"
)

func InsterOvirtDBConf(o *models.OvirtConf)(rowAffected int64,err error) {
	if rowAffected,err = mysql.InsertOvirtDBConf(o);err != nil{
		return 0,err
	}else{
		return rowAffected,nil
	}
}


func CheckOvirtDBConf(o *models.OvirtConf)(err error){
	if err = mysql.CheckOvirtDBConf(o);err != nil{
		return err
	}else{
		return nil
	}
}


func ListOvirtConf()(listovirt []*models.ListOvirtConf,err error) {
	listovirt, err = mysql.ListOvirt()
	if err != nil{
		//zap.L().Error("ListOvirtConf error!",zap.Error(err))
		return nil,err
	}
	return listovirt,nil
}

func GetOvirtConfDetail(aliasname string)(getovirtconfdetail []*models.GetOvirtConfDetail,err error){
	getovirtconfdetail,err = mysql.GetOvirtConfDetail(aliasname)
	if err != nil{
		return nil,err
	}
	return getovirtconfdetail,nil
}