package mysql

import (
	"errors"
	"example/fundemo01/web-app/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetCommunityList()([]*models.CommunityList,error){
	var communityList []*models.CommunityList
	result := Mysql.Find(&communityList)
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		zap.L().Warn("GetCommunityList Record Not Found!",zap.Error(result.Error))
		return nil,result.Error
	}
	return communityList,nil
	//rows,err := Mysql.Raw(`select community_id,community_name from community`).Rows()
	//if err != nil {
	//	return nil,err
	//}
	//defer rows.Close()
	//for rows.Next(){
	//	Mysql.ScanRows(rows,&communityList)
	//}
	//return &communityList,nil
}


func GetCommunityDetail(id int)([]*models.CommunityDetail,error){
	var communityDetail []*models.CommunityDetail
	//result := Mysql.Where("community_id = ?",id).Find(&communityDetail)
	result := Mysql.Raw("select community_id,community_name,introduction,create_time from community where community_id = ?",id).Scan(&communityDetail)
	if result.Error != nil{
		zap.L().Error("GetCommunityDetail Error!")
		return nil,result.Error
	}
	if result.RowsAffected == 0{
		zap.L().Warn("GetCommunityDetail Record Not Found!")
		return nil,errors.New("GetCommunityDetail Record Not Found!")
	}
	//if errors.Is(result.Error,gorm.ErrRecordNotFound){
	//	zap.L().Warn("GetCommunityDetail Record Not Found!",zap.Error(result.Error))
	//	return nil,result.Error
	//}
	return communityDetail,nil
	//rows,err := Mysql.Raw(`select community_id,community_name from community`).Rows()
	//if err != nil {
	//	return nil,err
	//}
	//defer rows.Close()
	//for rows.Next(){
	//	Mysql.ScanRows(rows,&communityList)
	//}
	//return &communityList,nil
}
