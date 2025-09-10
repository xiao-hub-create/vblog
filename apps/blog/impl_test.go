package blog_test

import (
	"context"
	"testing"

	"github.com/xiao-hub-create/vblog/apps/blog"
	"github.com/xiao-hub-create/vblog/apps/blog/impl"
)

var (
	ctx = context.Background()
)

func TestCreateBlog(t *testing.T) {
	req := &blog.CreateBlogRequest{
		Title:    "测试标题",
		Summary:  "测试摘要",
		Content:  "测试内容",
		Category: "测试分类",
		Tags: map[string]string{
			"gogogo": "1111",
		},
	}
	ins, err := impl.BlogService.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestListBlog(t *testing.T) {
	req := blog.NewListBlogRequset()
	req.Tags = map[string]string{
		"gogogo": "1111",
	}
	ins, err := impl.BlogService.ListBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestNewListBlogRequest(t *testing.T) {
	req := blog.NewListBlogRequset()
	req.SetTag("key1=value1,key2=value2")
	t.Log(req)
}
