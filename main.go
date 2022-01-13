package main

import (
	"hitszedu-go/config"
	"hitszedu-go/database"
	"hitszedu-go/router"
	"hitszedu-go/service"
)

func main() {
	config.Init()           //加载配置文件
	database.ConnectMysql() //连接数据库
	router.InitRouter()     //加载路由
	service.WxStart()       //开启微信服务
}
