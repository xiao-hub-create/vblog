package user

import (
	"time"

	"github.com/xiao-hub-create/vblog/utils"
)

type User struct {
	//存放到数据里的对象的元数据
	utils.ResourceMeta
	//具体参数
	RegistryRequest
}

type RegistryRequest struct {
	Username string `json:"username" gorm:"column:username;unique;index"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)"`
	//用户资料
	Profile
	//用户状态
	Status
}

type Profile struct {
	//头像
	Avatar string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	//昵称
	NickName string `json:"nic_name" gorm:"column:nic_name;type:varchar(100)"`
	//邮箱
	Email string `json:"email" gorm:"column:email;type:varchar(100)"`
}

type Status struct {
	//封禁时间
	BlockAt *time.Time `json:"block_at" gorm:"column:block_at"`
	//封禁原因
	BlockReason string `json:"block_reason" gorm:"column:block_reason;type:text"`
}

func (s *Status) IsBlock() bool {
	return s.BlockAt != nil
}

func (u User) TableName() string {
	return "users"
}
