package main

import (
	"fmt"

	"rongqin.cn/todo_task/conf"
	"rongqin.cn/todo_task/model"
	"rongqin.cn/todo_task/router"
)

var Config conf.AppConfig

func init() {
	conf.ConfigLoad()                              // 加载配置文件
	model.DataSourceMysql(*conf.Config.DataSource) // 加载数据库配置
	model.CreateTables()                           //创建表结构

}

func main() {

	router := router.ApiRouter()
	router.Run(conf.Config.HttpConfig.HttpPort)
	fmt.Println("应用启动完成.")
}
