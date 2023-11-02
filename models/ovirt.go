package models

type OvirtConf struct {
	AliasName         string `json:"aliasname" binding:"required"`
	Host              string `json:"host" binding:"required"`
	Port              int    `json:"port" binding:"required"`
	User              string `json:"user" binding:"required"`
	Password          string `json:"password" binding:"required"`
	DBName            string `json:"dbname" binding:"required"`
	TimeZone          string `json:"timezone" binding:"required"`
	PGMaxIdelConns    int    `json:"pgmaxidelconns" gorm:"default:10"`
	PGMaxOpenConns    int    `json:"pgmaxopenconns" gorm:"default:20"`
	PGConnMaxLifetime int    `json:"pgconnmaxlifetime" gorm:"default:15"`
	PGPreStatement    string `json:"pgprestatement" gorm:"default:'true'"`
}

type ListOvirtConf struct {
	AliasName string `json:"aliasname"  gorm:"column:aliasname"`
	Host      string `json:"host" gorm:"column:host"`
	Port      int    `json:"port" gorm:"column:port"`
}

type GetOvirtConfDetail struct {
	ID                int    `json:"id" gorm:"column:id"`
	AliasName         string `json:"aliasname" gorm:"column:aliasname"`
	Host              string `json:"host" gorm:"column:host"`
	Port              int    `json:"port" gorm:"column:port"`
	DBName            string `json:"dbname" gorm:"column:dbname"`
	TimeZone          string `json:"timezone" gorm:"column:timezone"`
	PGMaxIdelConns    int    `json:"pgmaxidelconns" gorm:"column:pgmaxidelconns"`
	PGMaxOpenConns    int    `json:"pgmaxopenconns" gorm:"column:pgmaxopenconns"`
	PGConnMaxLifetime int    `json:"pgconnmaxlifetime" gorm:"column:pgconnmaxlifetime"`
	PGPreStatement    string `json:"pgprestatement" gorm:"column:pgprestatement"`
}
