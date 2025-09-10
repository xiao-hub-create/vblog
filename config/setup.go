package config

import (
	"fmt"
	"os"

	"github.com/infraboard/mcube/v2/ioc"
)

func LoadConfig() error {
	//配置单元测试的配置
	req := ioc.NewLoadConfigRequest()
	// fmt.Printf("req: %+v", req)
	req.ConfigFile.Enabled = true
	// //必须配置绝对路径，{workspace}/etc/application.toml

	// os.Setenv("WORKSPACE_DIR", "/data/workspace/vblog")
	// fmt.Println(os.Getenv("WORKSPACE_DIR"))
	workspaceDir := os.Getenv("WORKSPACE_DIR")
	if workspaceDir == "" {
		req.ConfigFile.Path = "etc/application.toml"
	} else {
		req.ConfigFile.Path = workspaceDir + "/etc/application.toml"
	}
	// req.ConfigFile.Path = "/data/workspace/vblog/etc/application.toml"
	err := ioc.ConfigIocObject(req)
	if err != nil {
		return err
	}
	fmt.Printf("WORKSPACE_DIR: %s\n", req.ConfigFile.Path)
	return nil
}
