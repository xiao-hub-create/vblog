package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/xiao-hub-create/vblog/apps/user"
	"golang.org/x/crypto/bcrypt"
)

var UserService user.Service = &UserServiceImpl{}

func init() {
	ioc.Controller().Registry(&UserServiceImpl{})
}

var _ user.Service = (*UserServiceImpl)(nil)

// 用于实现UserService接口
type UserServiceImpl struct {
	// 继承ioc.ObjectImpl，实现了ioc.Object接口
	ioc.ObjectImpl
}

func (i *UserServiceImpl) Name() string {
	return user.AppName
}

func (u UserServiceImpl) UpdateUserStatus(ctx context.Context, req *user.UpdateUserStatusRequest) (*user.User, error) {
	panic("unimplemented")
}

func (u UserServiceImpl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	query := datasource.DBFromCtx(ctx)
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		query = query.Where("id = ?", req.Value)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("username = ?", req.Value)
	}
	ins := &user.User{}
	if err := query.Take(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (u UserServiceImpl) Registry(ctx context.Context, req *user.RegistryRequest) (*user.User, error) {
	ins, err := user.New(req)
	if err != nil {
		fmt.Println("create user failed", err)
		return nil, err
	}
	//明文密码保存到数据库不安全
	//对称加密/非对称 ，解密
	//消息摘要，无法还原
	hashPass, err := bcrypt.GenerateFromPassword([]byte(ins.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	ins.Password = string(hashPass)

	//无事务的模式
	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (u UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) error {
	panic("unimplemented")
}

func (u UserServiceImpl) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) error {
	panic("unimplemented")
}

func (u UserServiceImpl) UpdateProfile(ctx context.Context, req *user.UpdateProfileRequest) (*user.User, error) {
	panic("unimplemented")
}

func (u UserServiceImpl) Unregistry(ctx context.Context, req *user.UnregistryRequest) error {
	panic("unimplemented")
}
