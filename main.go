package main

import (
	"fmt"

	"rongqin.cn/todo_task/conf"
	"rongqin.cn/todo_task/model"
)

func init() {
	var config conf.AppConfig
	config.ConfigLoad()                       // 加载配置文件
	model.DataSourceMysql(*config.DataSource) // 加载数据库配置
	model.CreateTables()                      //创建表结构

}

func main() {

	fmt.Println("应用启动完成.")
}
