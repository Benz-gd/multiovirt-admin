package models

type ListHostGroup struct {
	GroupID     string `json:"groupid"  gorm:"column:groupid"`
	GroupName   string `josn:"groupname" gorm:"column:groupname"`
	Status      string `josn:"status" gorm:"column:status"`
	Description string `json:"description" gorm:"column:description"`
	CreateBy    string `json:"createby" gorm:"column:createby"`
	ModifiedBy  string `json:"modifiedby" gorm:"column:modifiedby"`
	CreatedOn   string `json:"description" gorm:"column:description"`
	ModifiedOn  string `json:"description" gorm:"column:description"`
}
