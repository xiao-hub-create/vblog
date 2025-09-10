package blog

import (
	"context"
	"strings"

	"github.com/xiao-hub-create/vblog/utils"
)

type Service interface {
	// 创建博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 博客列表查询
	ListBlog(context.Context, *ListBlogRequest) (*BlogSet, error)
	// 博客详情查询
	DetailBlog(context.Context, *DetailBlogRequest) (*Blog, error)
	// 博客编辑
	EditBlog(context.Context, *EditBlogRequest) (*Blog, error)
	// 发布
	PublishBlog(context.Context, *PublishBlogRequest) (*Blog, error)
	// 删除
	DeleteBlog(context.Context, *DetailBlogRequest) error
}

func NewListBlogRequset() *ListBlogRequest {
	return &ListBlogRequest{
		PageRequest: *utils.NewPageRequest(),
		Tags:        map[string]string{},
	}
}

type ListBlogRequest struct {
	//分页参数
	utils.PageRequest
	//关键字查询参数，标题，内容，摘要，作者，分类，标签
	Keywords string `json:"keywords" form:"keywords"`
	//状态过滤参数，作者：nil,访客：STAGE_PUBLISHED
	Stage *STAGE `json:"stage" form:"stage"`
	//查询某个用户具体的文章
	CreateBy string `json:"create_by" form:"create_by"`
	//查询分类的文章
	Category string `json:"category" form:"category"`
	//查询tag相关文章
	Tags map[string]string `json:"tags" form:"-"`
}

func (r *ListBlogRequest) SetTag(tags string) {
	kvItem := strings.Split(tags, ",")
	for i := range kvItem {
		kvString := kvItem[i]
		kv := strings.Split(kvString, "=")
		if len(kv) > 1 {
			r.Tags[kv[0]] = strings.Join(kv[1:], "=")
		}

	}
}

type DetailBlogRequest struct {
	//博客id
	utils.GetRequest
}

type EditBlogRequest struct {
	//博客id
	utils.GetRequest
	CreateBlogRequest
}

type PublishBlogRequest struct {
	//博客id
	utils.GetRequest
	StatusSpec
}

type DeleteBlogRequest struct {
	//博客id
	utils.GetRequest
}
