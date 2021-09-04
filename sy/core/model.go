package core

import (
	"sy/models"
	"time"
)

type Water struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Category  string    //1数据库 2file 3excel
	Content   string    //内容
	Report    string    //反馈结果
	Mean      float64   //中文均值
	Message   string    //备注
}

func init() {
	models.DB.AutoMigrate(&Water{})
}
