/*************************************************************************
	> File Name: Server.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/

package main

import (
	"fmt"
	"flag"
	"shortlib"
)

func main(){
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
	redis-cli,err := shortlib.NewRedisAdaptor(configure)
	if err != nil {
		fmt.Printf("[ERROR] Redis init fail..\n")
		return
	}
	
	//启动服务
	



}





