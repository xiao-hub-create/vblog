package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/xiao-hub-create/vblog/apps/blog"
	"github.com/xiao-hub-create/vblog/middleware"
)

var BlogService blog.Service = &BlogServiceImpl{}

type BlogServiceImpl struct {
}

func (b BlogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	ins, err := blog.NewBlog(req)
	if err != nil {
		return nil, err
	}

	tk := middleware.GetTokenFromContext(ctx)
	ins.CreateBy = tk.RefUserName

	err = datasource.DBFromCtx(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil

}

func (b BlogServiceImpl) ListBlog(ctx context.Context, req *blog.ListBlogRequest) (*blog.BlogSet, error) {
	query := datasource.DBFromCtx(ctx).Model(&blog.Blog{})
	if req.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	}
	if req.Stage != nil {
		query = query.Where("stage = ?", req.Stage)
	}
	if req.CreateBy != "" {
		query = query.Where("create_by = ?", req.CreateBy)
	}
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}
	for k, v := range req.Tags {
		query = query.Where(fmt.Sprintf("JSON_UNQUOTE(JSON_EXTRACT(tags, '$.%s')) = ?", k), v)
	}

	set := blog.NewBlogSet()

	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	err := query.Order("created_at DESC").Offset(int(req.Offset())).Limit(int(req.PageSize)).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (b BlogServiceImpl) DetailBlog(ctx context.Context, req *blog.DetailBlogRequest) (*blog.Blog, error) {
	panic("unimplemented")
}

func (b BlogServiceImpl) EditBlog(ctx context.Context, req *blog.EditBlogRequest) (*blog.Blog, error) {
	panic("unimplemented")
}

func (b BlogServiceImpl) PublishBlog(ctx context.Context, req *blog.PublishBlogRequest) (*blog.Blog, error) {
	panic("unimplemented")
}

func (b BlogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DetailBlogRequest) error {
	panic("unimplemented")
}
