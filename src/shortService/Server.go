/*************************************************************************
	> File Name: Server.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"shortlib"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "config.ini", "configure file full path")
	flag.Parse()

	//读取配置文件
	configure, err := shortlib.NewConfigure(configFile)
	if err != nil {
		fmt.Printf("[ERROR] Parse Configure File Error: %v\n", err)
		return
	}

	//启动Redis客户端
	redis_cli, err := shortlib.NewRedisAdaptor(configure)
	if err != nil {
		fmt.Printf("[ERROR] Redis init fail..\n")
		return
	}
	if configure.GetRedisStatus() {
		err = redis_cli.InitCountService()
		if err != nil {
			fmt.Printf("[ERROR] Init Redis key count fail...\n")
		}
	}

	//不使用redis的情况下，启动短链接计数器
	count_channl := make(chan CountChannl, 1000)
	go CountThread(count_channl)

	//启动LRU缓存
	lru,err := shortlib.NewLRU(redis_cli)
	if err != nil {
		fmt.Printf("[ERROR]LRU init fail...\n")
	}
	//初始化两个短连接服务
	baseprocessor := &shortlib.BaseProcessor{redis_cli, configure, configure.GetHostInfo(),lru}

	original := &OriginalProcessor{baseprocessor, count_channl}
	short := &ShortProcessor{baseprocessor}

	//启动http handler
	router := &shortlib.Router{configure, map[int]shortlib.Processor{
		0: short,
		1: original,
	}}

	//启动服务

	port, _ := configure.GetPort()
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("[INFO]服务启动。。。地址:%v,端口:%v\n", addr, port)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		//logger.Error("Server start fail: %v", err)
		os.Exit(1)
	}

}
