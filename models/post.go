package models

type ParamPost struct {
	Post_ID int64 `json:"post_id" gorm:"column:post_id"`
	Title string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

