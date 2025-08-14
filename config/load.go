package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var conf = &Config{
	App: &App{
		Host: "0.0.0.0",
		Port: 8080,
	},
	MySQL: &MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
		Username: "root",
		Password: "123456789",
		Debug:    true,
	},
	Log: &Log{},
}

func Get() *Config {
	if conf == nil {
		panic("配置未初始化")
	}
	return conf
}

//配置的加载

// 程序其他部分如何读取程序配置
// yaml file -->config
func LoadConfigFromYaml(configFilePath string) error {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	//默认值 defaultConfig
	//结合配置文件 传达进来的参数
	//default <-- user define

	return yaml.Unmarshal(content, conf)
}
