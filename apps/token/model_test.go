package token_test

import (
	"testing"

	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/xiao-hub-create/vblog/apps/token"
	"github.com/xiao-hub-create/vblog/config"
)

func TestMigarate(t *testing.T) {
	if err := datasource.DB().AutoMigrate(&token.Token{}); err != nil {
		t.Fatal(err)
	}
}

func init() {
	config.LoadConfig()
}
