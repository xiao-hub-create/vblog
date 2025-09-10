package user

import (
	"time"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/xiao-hub-create/vblog/utils"
	"golang.org/x/crypto/bcrypt"
)

func New(req *RegistryRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("参数校验失败: %s", err)
	}
	return &User{
		ResourceMeta:    *utils.NewResourceMeta(),
		RegistryRequest: *req,
	}, nil
}

type User struct {
	//存放到数据里的对象的元数据
	utils.ResourceMeta
	//具体参数
	RegistryRequest
}

func (r User) String() string {
	return pretty.ToJSON(r)
}
func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{}
}

type RegistryRequest struct {
	Username string `json:"username" gorm:"column:username;unique;index" validate:"required"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)" validate:"required"`
	//用户资料
	Profile
	//用户状态
	Status
}

func (r *RegistryRequest) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(password))
}

func (r *RegistryRequest) Validate() error {
	return validator.Validate(r)
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
