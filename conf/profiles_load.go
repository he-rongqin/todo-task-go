package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// 应用上下文配置
type AppContextConfig struct {
	Active string `ini:"app.active"` // 当前激活的配置
	Mode   string `ini:"app.mode"`   //应用模式，dev、test、prod

}

// 应用配置
type ServerConfig struct {
	HttpPort    string `ini:"server.port"` // 应用http 端口，默认：8080
	ContextPath string `ini:"server.path"` // http访问根路径
}

// mysql数据源配置
type DataSourceMysql struct {
	DSN          string `ini:"ds.mysql.dsn"`          // 数据库地址
	MaxIdleConns int    `ini:"ds.mysql.MaxIdleConns"` // 最大连接数
	MaxOpenConns int    `int:"ds.mysql.MaxOpenConns"` //设置打开数据库连接的最大数量
}

type AppConfig struct {
	Context    *AppContextConfig // 上下文配置
	HttpConfig *ServerConfig     // http 服务配置
	DataSource *DataSourceMysql  // 数据库配置
}

// 加载ini 配置文件
func (config *AppConfig) ConfigLoad() {
	fmt.Println("正在加载配置文件...")
	cfg, err := ini.Load("./conf/config.ini") // load 配置文件
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		panic("加载配置文件异常。")
	}

	// 加载上下文配置
	context := new(AppContextConfig)
	cfg.Section("app.context").MapTo(context)
	fmt.Printf("context init : %v\n", context)
	config.Context = context
	// 加载http 配置
	h := new(ServerConfig)
	cfg.Section("http.server").MapTo(h)
	fmt.Printf("server config init : %v\n", h)
	config.HttpConfig = h
	// 加载数据库配置
	active := context.Active
	ds := new(DataSourceMysql)
	cfg.Section("active." + active).MapTo(ds)
	fmt.Printf("datasource for mysql config init: %v\n", ds)
	config.DataSource = ds

}
