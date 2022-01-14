package controller

import (
	"ctjsoft/ginessential/common"
	"ctjsoft/ginessential/model"
	"ctjsoft/ginessential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)

	// 判断手机号是否存在
	if isPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

// 判断手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
