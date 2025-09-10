package token

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc/config/validator"
)

// 业务域
type Service interface {
	UserService
	InnerService
}

// 外部接口
type UserService interface {
	//颁发令牌 登录
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)

	//注销令牌 退出
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

type IssueTokenRequest struct {
	//用户名
	Username string `json:"username" validate:"required"`
	//密码
	Password string `json:"password" validate:"required"`
	//Token可能1天过期，过期时间调整为7天
	RememberMe bool `json:"remember_me"`
}

func (r *IssueTokenRequest) Validate() error {
	return validator.Validate(r)
}

type RevokeTokenRequest struct {
	//访问Token
	AccessToken string `json:"access_token"`
	//刷新Token
	//与访问令牌构成一对，避免AccessToken泄露
	RefreshToken string `json:"refresh_token"`
}

// 内部接口
type InnerService interface {
	//校验令牌
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

type ValidateTokenRequest struct {
	//访问令牌
	AccessToken string `json:"access_token"`
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}
