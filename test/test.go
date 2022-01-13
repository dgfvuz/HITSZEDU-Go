package main

import (
	"fmt"
	"hitszedu-go/config"
	"hitszedu-go/database"
	"hitszedu-go/model"
)

func main() {
	config.Init()           //加载配置文件
	database.ConnectMysql() //连接数据库
	user := model.GetUser("c716bf7c-58af-43a9-8c57-c237278fe9e4")
	fmt.Println(user)
}
