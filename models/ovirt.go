package models

type OvirtConf struct {
	AliasName string `json:"aliasname" binding:"required"`
	Host string `josn:"host" binding:"required"`
	Port int `josn:"port" binding:"required"`
	User string `josn:"user" binding:"required"`
	Password string `josn:"password" binding:"required"`
	DBName string `josn:"dbname" binding:"required"`
	TimeZone string `josn:"timezone" binding:"required"`
	PGMaxIdelConns int `josn:"pgmaxidelconns" gorm:"default:10"`
	PGMaxOpenConns int `josn:"pgmaxopenconns" gorm:"default:20"`
	PGConnMaxLifetime int `josn:"pgconnmaxlifetime" gorm:"default:15"`
	PGPreStatement string `josn:"pgprestatement" gorm:"default:'true'"`
}

type ListOvirtConf struct {
	AliasName string `json:"aliasname"  gorm:"column:aliasname"`
	Host string `josn:"host" gorm:"column:host"`
	Port int `josn:"port" gorm:"column:port"`
}

type GetOvirtConfDetail struct {
	ID int `json:"id" gorm:"column:id"`
	AliasName string `json:"aliasname" gorm:"column:aliasname"`
	Host string `josn:"host" gorm:"column:host"`
	Port int `josn:"port" gorm:"column:port"`
	DBName string `josn:"dbname" gorm:"column:dbname"`
	TimeZone string `josn:"timezone" gorm:"column:timezone"`
	PGMaxIdelConns int `josn:"pgmaxidelconns" gorm:"column:pgmaxidelconns"`
	PGMaxOpenConns int `josn:"pgmaxopenconns" gorm:"column:pgmaxopenconns"`
	PGConnMaxLifetime int `josn:"pgconnmaxlifetime" gorm:"column:pgconnmaxlifetime"`
	PGPreStatement string `josn:"pgprestatement" gorm:"column:pgprestatement"`
}