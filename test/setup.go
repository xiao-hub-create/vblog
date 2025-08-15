package test

import "github.com/infraboard/mcube/v2/ioc"

func DevelopmentSetup() {
	//配置单元测试的配置
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	//必须配置绝对路径，{workplace}/etc/application.toml
	req.ConfigFile.Path = ""
	err := ioc.ConfigIocObject(req)
	if err != nil {
		panic(err)
	}

}
