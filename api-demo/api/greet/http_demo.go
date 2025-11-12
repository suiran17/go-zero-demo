package main

import (
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

/*
运行:

	$ go run http_demo.go
	{"@timestamp":"2025-06-22T15:24:18.433+08:00","caller":"stat/usage.go:61","content":"CPU: 0m, MEMORY: Alloc=1.8Mi, TotalAlloc=1.8Mi, Sys=12.7Mi, NumGC=0","level":"stat"}
	{"@timestamp":"2025-06-22T15:24:18.433+08:00","caller":"load/sheddingstat.go:61","content":"(api) shedding_stat [1m], cpu: 0, total: 0, pass: 0, drop: 0","level":"stat"}
	{"@timestamp":"2025-06-22T15:24:38.834+08:00","caller":"handler/loghandler.go:147","content":"[HTTP] 200 - GET /hello/world - 127.0.0.1:59378 - curl/8.7.1","duration":"0.1ms","level":"info","span":"0c928bf2570cace8","trace":"6418948df54d41af79c2e7bb6e850a0a"}

测试:

	$ curl localhost:8080/hello/world
	"Hello World!"
*/
func main() {
	var restConf rest.RestConf
	conf.MustLoad("etc/helloworld.yaml", &restConf)
	s, err := rest.NewServer(restConf)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.AddRoute(rest.Route{ // 添加路由
		Method: http.MethodGet,
		Path:   "/hello/world",
		Handler: func(writer http.ResponseWriter, request *http.Request) { // 处理函数
			httpx.OkJson(writer, "Hello World!")
		},
	})

	defer s.Stop()
	s.Start() // 启动服务
}
