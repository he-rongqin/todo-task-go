package model

// 创建数据库表结构
func CreateTables() {

	// 给表添加后缀
	dbMysql.Set("gorm:table_options", "ENGINE=InnoDB")
	// 表存在时，不做新创建，而是更新
	dbMysql.AutoMigrate(&User{})

}
