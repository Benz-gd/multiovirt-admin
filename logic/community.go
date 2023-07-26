package logic

import (
	"example/fundemo01/web-app/dao/mysql"
	"example/fundemo01/web-app/models"
)

func GetCommunityList()([]*models.CommunityList,error) {
     //查找数据库，查找所有的community，并返回
	var data []*models.CommunityList
	var err error
	data,err = mysql.GetCommunityList()
	if err != nil {
		return nil,err
	}
	return data,nil
}


func GetCommunityDetail(id int)([]*models.CommunityDetail,error) {
	//查找数据库，查找所有的community，并返回
	var data []*models.CommunityDetail
	var err error
	data,err = mysql.GetCommunityDetail(id)
	if err != nil {
		return nil,err
	}
	return data,nil
}
