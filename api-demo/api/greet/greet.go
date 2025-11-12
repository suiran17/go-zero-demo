package main

import (
	"flag"
	"fmt"

	"greet/internal/config"
	"greet/internal/handler"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

// main函数是程序的入口点。
func main() {
	// 解析命令行参数。
	flag.Parse()

	// 初始化配置对象。
	var c config.Config

	// 加载配置文件，如果加载失败，程序将panic。
	conf.MustLoad(*configFile, &c)

	// 创建并启动一个新的服务器实例。
	server := rest.MustNewServer(c.RestConf)
	// 确保在程序退出时停止服务器。
	defer server.Stop()

	// 创建服务上下文，包含业务逻辑所需的数据。
	ctx := svc.NewServiceContext(c)
	// 注册所有HTTP处理函数。
	handler.RegisterHandlers(server, ctx)

	// 启动服务器前，打印服务器启动信息。
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 启动服务器并开始监听。
	server.Start()
}
