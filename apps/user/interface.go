package user

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
)

// 获取实现，从ioc
func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

const (
	AppName = "user"
)

type Service interface {
	AdminService
	UserService
}

type AdminService interface {
	//更新用户状态
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*User, error)
	//查询某个具体的用户
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

type UserService interface {
	//注册
	Registry(context.Context, *RegistryRequest) (*User, error)

	//更新
	//密码更新
	UpdatePassword(context.Context, *UpdatePasswordRequest) error
	//忘记密码，密码重置，安全等级高
	ResetPassword(context.Context, *ResetPasswordRequest) error
	//更新用户资料
	UpdateProfile(context.Context, *UpdateProfileRequest) (*User, error)

	//注销
	Unregistry(context.Context, *UnregistryRequest) error
}

type UpdateUserStatusRequest struct {
	UserId string `json:"user_id"`
	Status
}

type UpdatePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordRequest struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
	//验证码
	VerifyCode string `json:"verify_code"`
}

type UpdateProfileRequest struct {
	//用户id
	UserId string `json:"user_id"`
	Profile
}

type UnregistryRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 通过username或者id查询用户
type DescribeUserRequest struct {
	DescribeBy DESCRIBE_BY `json:"describe_by"`
	Value      string      `json:"value"`
}
