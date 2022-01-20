package controller

import (
	"ctjsoft/ginessential/common"
	"ctjsoft/ginessential/dto"
	"ctjsoft/ginessential/model"
	"ctjsoft/ginessential/reponse"
	"ctjsoft/ginessential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(phone) != 11 {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为 11 位")
		return
	}
	if len(password) < 6 {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于 6 位")
		return
	}
	// 如果没有传名称, 给一个 10 位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)

	// 判断手机号是否存在
	if isPhoneExist(DB, phone) {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 加密
	if err != nil {
		reponse.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})
	reponse.Success(ctx, nil, "注册成功")
}

// Login 登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为 11 位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为 11 位"})
		return
	}
	if len(password) < 6 {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于 6 位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		reponse.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		reponse.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Panicf("token generate error: %v", err)
		return
	}

	// 返回结果
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})
	reponse.Success(ctx, gin.H{"token": token}, "登录成功")
}

// Info 获取用户信息
func Info(ctx *gin.Context) {
	// 获取用户信息时, 用户应该是通过认证的, 我们应该能从上下文获取到用户的信息
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

// 判断手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
