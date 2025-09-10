package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/xiao-hub-create/vblog/apps/token"
	"github.com/xiao-hub-create/vblog/apps/token/impl"
)

// 补充鉴权逻辑
// 1. 首先要获取Token, Cookie, Header, Authorization: 用于存放用户认证信息, Authorization: <token_type> <token_value>, Bearer xxxxxx
// 2. 校验Token
// 3. 注入用户信息
func Auth(c *gin.Context) {
	//1. 获取Token
	accessToken := c.GetHeader("Authorization")
	tklist := strings.Split(accessToken, " ")

	accessToken = ""
	if len(tklist) == 2 {
		accessToken = tklist[1]
	}

	//2. 校验Token
	tk, err := impl.TokenService.ValidateToken(c.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		response.Failed(c, exception.NewUnauthorized("令牌校验失败:%s", err))
		c.Abort()
		return
	}

	//	3. 注入用户信息
	ctx := context.WithValue(c.Request.Context(), TokenCtxKey{}, tk)
	c.Request = c.Request.WithContext(ctx)
}

type TokenCtxKey struct {
}

func GetTokenFromContext(ctx context.Context) *token.Token {
	tk, ok := ctx.Value(TokenCtxKey{}).(*token.Token)
	if !ok {
		return nil
	}
	return tk
}
