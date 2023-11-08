package mysql

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"multiovirt-admin/models"
)

func GetCommunityList() ([]*models.CommunityList, error) {
	var communityList []*models.CommunityList
	result := MysqlBase.Find(&communityList)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		zap.L().Warn("GetCommunityList Record Not Found!", zap.Error(result.Error))
		return nil, result.Error
	}
	return communityList, nil
}

func GetCommunityDetail(id int) ([]*models.CommunityDetail, error) {
	var communityDetail []*models.CommunityDetail
	result := MysqlBase.Raw("select community_id,community_name,introduction,create_time from community where community_id = ?", id).Scan(&communityDetail)
	if result.Error != nil {
		zap.L().Error("GetCommunityDetail Error!")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("GetCommunityDetail Record Not Found!")
		return nil, errors.New("GetCommunityDetail Record Not Found!")
	}
	return communityDetail, nil
}
