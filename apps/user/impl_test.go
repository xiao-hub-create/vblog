package user_test

import (
	"context"
	"testing"

	"github.com/xiao-hub-create/vblog/apps/user"
)

var (
	ctx = context.Background()
)

func TestRegistry(t *testing.T) {

	req := user.NewRegistryRequest()

	req.Username = "admin"
	req.Password = "123456"

	ins, err := user.GetService().Registry(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribe(t *testing.T) {
	ins, err := user.GetService().DescribeUser(ctx, &user.DescribeUserRequest{
		user.DESCRIBE_BY_USERNAME, "admin",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.CheckPassword("123456"))
}
