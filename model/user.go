package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Company  string
	Status   int8
}

// 用户名是否存在
func (user *User) Exist() bool {
	var count int64
	dbMysql.Model(&User{}).Where(&User{Username: user.Username, Company: user.Company}).Count(&count)
	return count > 0
}

// 新增数据前勾子，加密密码
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 密码加密
	pwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwd)
	return
}

// 新增数据
func (user *User) Insert() uint {
	d := dbMysql.Create(&user)
	if d.Error != nil {
		dbMysql.Logger.Error(dbMysql.Statement.Context, d.Error.Error())
	}
	return user.ID
}

// 根据用户名查找
func (user *User) GetUserByName(username string) {
	dbMysql.Where(&User{Username: username}).First(&user)

}

// 用户详情
func (user *User) Get(uid int) *User {
	dbMysql.First(&user, uid)
	return user
}
