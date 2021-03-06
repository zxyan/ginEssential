package common

import (
	"ctjsoft/ginessential/model"
	"fmt"
	"github.com/spf13/viper"
	"net/url"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB 连接数据库
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName") // "mysql"
	host := viper.GetString("datasource.host")             // "localhost"
	port := viper.GetString("datasource.port")             // "3306"
	database := viper.GetString("datasource.database")     // "ginessential"
	username := viper.GetString("datasource.username")     // "root"
	password := viper.GetString("datasource.password")     // "root"
	charset := viper.GetString("datasource.charset")       // "utf8"
	loc := viper.GetString("datasource.loc")               // 时区
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc), // 需要将 url encode 一下
	)
	// database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动迁移
	// db.AutoMigrate(&model.User{})
	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
		if db.HasTable(&model.User{}) {
			fmt.Println("User 表创建成功")
		} else {
			fmt.Println("User 表创建失败")
		}
	} else {
		fmt.Println("User 表已存在")
	}

	DB = db
	return db
}

// GetDB 获取 DB 实例
func GetDB() *gorm.DB {
	return DB
}
