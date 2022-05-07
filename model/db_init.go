package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rongqin.cn/todo_task/conf"
)

var dbMysql *gorm.DB

func DataSourceMysql(dbconfig conf.DataSourceMysql) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbconfig.DSN, // DSN data source name
		DefaultStringSize:         256,          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	if dbconfig.MaxIdleConns == 0 {
		// 默认10
		dbconfig.MaxIdleConns = 10
	}
	sqlDB.SetMaxIdleConns(dbconfig.MaxIdleConns)

	if dbconfig.MaxOpenConns == 0 {
		// 默认100
		dbconfig.MaxOpenConns = 100
	}
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(dbconfig.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Printf("数据库连接错误 : %v\n", err)
		panic("数据库连接错误")
	}
	fmt.Printf("db 连接成功: %v\n", db)
	dbMysql = db
}
