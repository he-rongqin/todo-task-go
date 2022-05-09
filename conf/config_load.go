package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// config 全局变量
var Config AppConfig

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

// 鉴权配置
type SecurityConfig struct {
	PrivateKey string `ini:"security.jwt.PrivateKey"` // jwt 加密私钥
	ExpireTime uint   `ini:"security.jwt.ExpireTime"` //jwt 过期时间
}
type AppConfig struct {
	Context        *AppContextConfig // 上下文配置
	HttpConfig     *ServerConfig     // http 服务配置
	DataSource     *DataSourceMysql  // 数据库配置
	SecurityConfig *SecurityConfig   //安全配置
}

// 加载ini 配置文件
func ConfigLoad() {
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
	Config.Context = context
	// 加载环境配置
	loadEnvConfig(context.Active)
	fmt.Println("配置文件加载完成。")

}

// 加载环境配置
func loadEnvConfig(active string) {
	fmt.Println("正在加载配置文件...")
	path := "./conf/env/config-" + active + ".ini"
	cfg, err := ini.Load(path) // load 配置文件
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		panic("加载配置文件异常。")
	}
	// 加载http 配置
	h := new(ServerConfig)
	cfg.Section("http.server").MapTo(h)
	fmt.Printf("server config init : %v\n", h)
	Config.HttpConfig = h
	// 加载数据库配置
	ds := new(DataSourceMysql)
	cfg.Section("ds.mysql").MapTo(ds)
	fmt.Printf("datasource for mysql config init: %v\n", ds)
	Config.DataSource = ds
	// 加载安全配置
	security := new(SecurityConfig)
	cfg.Section("security").MapTo(security)
	fmt.Printf("security config load: %v\n", security)
	Config.SecurityConfig = security
}
