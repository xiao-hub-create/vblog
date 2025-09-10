package blog

import (
	"time"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/xiao-hub-create/vblog/utils"
)

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

func NewBlog(in *CreateBlogRequest) (*Blog, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("参数异常：%s", err)
	}
	return &Blog{
		ResourceMeta:      *utils.NewResourceMeta(),
		CreateBlogRequest: *in,
	}, nil

}

type Blog struct {
	utils.ResourceMeta
	CreateBlogRequest
	Status
}

func (b *Blog) String() string {
	return pretty.ToJSON(b)
}

func (b *Blog) TableName() string {
	return "blog"
}

type CreateBlogRequest struct {

	//标题
	Title string `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	//摘要
	Summary string `json:"summary" gorm:"column:summary;type:text" validate:"required"`
	//内容
	Content string `json:"content" gorm:"column:content;type:text" validate:"required"`
	//分类
	Category string `json:"category" gorm:"column:category;type:varchar(200);index" validate:"required"`
	//标签
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (r *CreateBlogRequest) Validate() error {
	return validator.Validate(r)
}

type Status struct {
	StatusSpec
	//状态变更时间
	ChangeAt *time.Time `json:"change_at" gorm:"column:change_at"`
}

type StatusSpec struct {
	//0:草稿 1:发布 2:审核...
	Stage int `json:"stage" gorm:"column:stage;type:tinyint(1);index"`
}
