package models

type Ping_IP struct {
	StartIP string `josn:"startip" binding:"required"`
	EndIP   string `josn:"endip" binding:"required"`
}
