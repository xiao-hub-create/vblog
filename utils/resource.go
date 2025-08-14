package utils

import "time"

type ResourceMeta struct {
	Id        uint      `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
	//用户
	CreateBy  string     `json:"create_by" gorm:"column:create_by;type:varchar(200)"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
