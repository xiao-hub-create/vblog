package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"

	"github.com/xiao-hub-create/vblog/apps/token"
	"github.com/xiao-hub-create/vblog/apps/user"
	"github.com/xiao-hub-create/vblog/apps/user/impl"
)

var TokenService token.Service = &TokenServiceImpl{
	user: impl.UserService,
}

type TokenServiceImpl struct {
	user user.AdminService
}

func (t *TokenServiceImpl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("参数校验异常：%s", err)
	}
	//1.查询用户
	u, err := t.user.DescribeUser(ctx, &user.DescribeUserRequest{
		DescribeBy: user.DESCRIBE_BY_USERNAME,
		Value:      req.Username,
	})
	if err != nil {
		return nil, err
	}
	//2.比对密码
	if err := u.CheckPassword(req.Password); err != nil {
		return nil, err
	}
	//3.生成token
	tk := token.NewToken(u.Id).SetRefUserName(u.Username)
	if err := datasource.DBFromCtx(ctx).Create(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

func (t *TokenServiceImpl) RevokeToken(ctx context.Context, req *token.RevokeTokenRequest) (*token.Token, error) {
	panic("not implemented")
}

// 1.查询token
// 2.token是否过期
func (t *TokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	tk := &token.Token{}
	if err := datasource.DBFromCtx(ctx).Where("access_token = ?", req.AccessToken).Take(tk).Error; err != nil {

		return nil, err
	}
	if err := tk.IsAccessTokenExpired(); err != nil {
		return nil, err
	}

	u, err := t.user.DescribeUser(ctx, &user.DescribeUserRequest{
		DescribeBy: user.DESCRIBE_BY_ID,
		Value:      tk.RefUserId,
	})
	if err != nil {
		return nil, err
	}

	tk.RefUserName = u.Username

	return tk, nil
}
