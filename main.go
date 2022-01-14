package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`    // 设置用户名(name)不为空
	Phone    string `gorm:"varchar(110);not null;unique"` // 设置电话号码(phone)唯一并且不为空
	Password string `gorm:"size:255;not null"`            // 设置字段大小为255
}

func main() {
	db := InitDB()
	defer db.Close()

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了 request 和 response
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(phone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为 11 位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于 6 位"})
			return
		}
		// 如果没有传名称, 给一个 10 位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, phone, password)

		// 判断手机号是否存在
		if isPhoneExist(db, phone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		// 创建用户
		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newUser)

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	// 3.监听端口，默认在 8080
	// Run("里面不指定端口号默认为 8080")
	panic(r.Run(":8000"))
}

// 随机生成字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiop")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 连接数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	// database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动迁移
	// db.AutoMigrate(&User{})
	if !db.HasTable(&User{}) {
		db.AutoMigrate(&User{})
		if db.HasTable(&User{}) {
			fmt.Println("User 表创建成功")
		} else {
			fmt.Println("User 表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}

	return db
}

// 判断手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
