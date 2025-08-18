package test

import (
	"fmt"
	"log"
	"os"

	"github.com/infraboard/mcube/v2/ioc"
)

func DevelopmentSetup() {
	//配置单元测试的配置
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	//必须配置绝对路径，{workspace}/etc/application.toml
	fmt.Println("=== 环境变量列表 ===")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	value := os.Getenv("WORKSPACE_DIR")
	fmt.Printf("WORKSPACE_DIR=%s\n", value)
	if value == "" {
		// 环境变量未设置或为空
		log.Fatal("WORKSPACE_DIR 环境变量未设置")
		// 或提供默认值：
		// value = "/default/path"
	}
	req.ConfigFile.Path = os.Getenv("WORKSPACE_DIR") + "/etc/application.toml"
	// req.ConfigFile.Path = "/data/workspace/vblog/etc/application.toml"
	err := ioc.ConfigIocObject(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("WORKSPACE_DIR: %s\n", req.ConfigFile.Path)
}
