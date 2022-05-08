package main

import (
	"fmt"

	"rongqin.cn/todo_task/conf"
	"rongqin.cn/todo_task/model"
	"rongqin.cn/todo_task/router"
)

var config conf.AppConfig

func init() {
	config.ConfigLoad()                       // 加载配置文件
	model.DataSourceMysql(*config.DataSource) // 加载数据库配置
	model.CreateTables()                      //创建表结构

}

func main() {

	router := router.ApiRouter()
	router.Run(config.HttpConfig.HttpPort)
	fmt.Println("应用启动完成.")
}
