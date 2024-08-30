package main

import (
	"flag"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/install"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/middleware"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

func main() {

	port := flag.Int("port", 3000, "port to listen on")
	username := flag.String("user", "admin", "admin username")
	password := flag.String("pass", "123456@2024", "admin password")
	flag.Parse()

	// 配置资源
	config := &builder.Config{

		// JWT加密密串
		AppKey: "123456789@#2024",

		// 加载服务
		Providers: service.Providers,

		// 数据库配置
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open("./data.db"),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// WEB根目录
	b.Static("/", "./web/app")

	// 自动构建数据库、拉取静态文件
	install.Handle(*username, *password)

	// 后台中间件
	b.Use(middleware.Handle)

	// 响应Get请求
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 启动服务
	b.Run(":" + fmt.Sprintf("%d", *port))
}
