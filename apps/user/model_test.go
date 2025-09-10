package user_test

import (
	"log"
	"testing"

	"github.com/xiao-hub-create/vblog/apps/user"
	"github.com/xiao-hub-create/vblog/config"

	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func TestMigrate(t *testing.T) {

	if err := Init(); err != nil {
		log.Fatalf("Could not initialize: %v", err)
	}
	//连接数据库
	if err := datasource.DB().AutoMigrate(&user.User{}); err != nil {
		t.Fatal(err)
	}

}

func Init() error {
	if err := config.LoadConfig(); err != nil {
		return err
	}
	return nil
}

//版本1
// func TestMigrate(t *testing.T) {
// 	path := os.Getenv("CONFIG_PATH")
// 	if path == "" {
// 		path = "../../config/application.yaml"
// 	}
// 	if err := config.LoadConfigFromYaml(path); err != nil {
// 		t.Fatalf("加载配置文件失败:%s", err)
// 	}

// 	//访问加载后的配置
// 	db := config.Get().MySQL.DB()

// 	if err := db.AutoMigrate(&user.User{}); err != nil {
// 		t.Fatal(err)
// 	}
// }
