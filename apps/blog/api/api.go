package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/xiao-hub-create/vblog/apps/blog"
	"github.com/xiao-hub-create/vblog/apps/blog/impl"
	"github.com/xiao-hub-create/vblog/middleware"
)

func NewBlogApiHandler() *BlogApiHandler {
	return &BlogApiHandler{
		blog: impl.BlogService,
	}
}

type BlogApiHandler struct {
	blog blog.Service
}

func (h *BlogApiHandler) Registry(r *gin.Engine) {
	router := r.Group("/vblog/api/v1/blogs")
	router.Use(middleware.Auth)
	router.POST("", h.CreateBlog)
	router.GET("", h.ListBlog)

}

func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	in := &blog.CreateBlogRequest{}
	if err := ctx.BindJSON(in); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := h.blog.CreateBlog(ctx.Request.Context(), in)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)

}

func (h *BlogApiHandler) ListBlog(ctx *gin.Context) {
	in := blog.NewListBlogRequset()
	if err := ctx.BindQuery(in); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := h.blog.ListBlog(ctx.Request.Context(), in)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}
