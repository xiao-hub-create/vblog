package blog_test

import (
	"testing"

	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/xiao-hub-create/vblog/apps/blog"
	"github.com/xiao-hub-create/vblog/config"
)

func TestMigarate(t *testing.T) {
	if err := datasource.DB().AutoMigrate(&blog.Blog{}); err != nil {
		t.Fatal(err)
	}
}

func init() {
	config.LoadConfig()
}
