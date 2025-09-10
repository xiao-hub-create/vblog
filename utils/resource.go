package utils

import "time"

func NewResourceMeta() *ResourceMeta {
	return &ResourceMeta{
		CreatedAt: time.Now(),
	}
}

type ResourceMeta struct {
	Id        string    `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
	//用户
	CreateBy  string     `json:"create_by" gorm:"column:create_by;type:varchar(200)"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
