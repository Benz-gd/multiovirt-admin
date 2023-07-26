package models

import "time"

type CommunityList struct {
	ID int64 `json:"id" gorm:"column:community_id"`
	Name string `json:"name" gorm:"column:community_name"`
}



type CommunityDetail struct {
	ID int64 `json:"id" gorm:"column:community_id"`
	Name string `json:"name" gorm:"column:community_name"`
	Introduction string `json:"introduction" gorm:"column:introduction"`
	Create_time time.Time  `json:"create_time" gorm:"column:create_time"`
}