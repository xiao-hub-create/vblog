package token_test

import (
	"context"
	"testing"

	"github.com/xiao-hub-create/vblog/apps/token"
	"github.com/xiao-hub-create/vblog/apps/token/impl"
)

var (
	ctx = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")

	ins, err := impl.TokenService.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("51bf49f5-12a2-406a-baf8-3f99d985b41a")
	ins, err := impl.TokenService.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
