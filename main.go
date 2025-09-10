package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/infraboard/mcube/v2/ioc/config/http"
	blogapi "github.com/xiao-hub-create/vblog/apps/blog/api"
	tokenapi "github.com/xiao-hub-create/vblog/apps/token/api"
	"github.com/xiao-hub-create/vblog/config"
)

func main() {
	config.LoadConfig()

	//gin Engine
	server := gin.Default()

	//注册业务模块的路由
	tokenapi.NewTokenApiHandler().Registry(server)
	blogapi.NewBlogApiHandler().Registry(server)

	//服务器启动
	if err := server.Run(http.Get().Addr()); err != nil {
		log.Println(err)
	}
}
