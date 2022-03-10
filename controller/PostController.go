package controller

import (
	"ctjsoft/ginessential/common"
	"ctjsoft/ginessential/model"
	"ctjsoft/ginessential/reponse"
	"ctjsoft/ginessential/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(&model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		reponse.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取登录用户 user
	user, _ := ctx.Get("user")

	// 创建文章 post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}

	reponse.Success(ctx, gin.H{"post": post}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		reponse.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取 path 中的参数
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		reponse.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		reponse.Fail(ctx, nil, "文章不属于您, 请勿非法操作")
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		reponse.Fail(ctx, nil, "更新失败")
		return
	}

	reponse.Success(ctx, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	// 获取 path 中的参数
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		reponse.Fail(ctx, nil, "文章不存在")
		return
	}

	reponse.Success(ctx, gin.H{"post": post}, "查看成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取 path 中的参数
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		reponse.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户 user
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		reponse.Fail(ctx, nil, "文章不属于您, 请勿非法操作")
		return
	}

	p.DB.Delete(&post)

	reponse.Success(ctx, gin.H{"post": post}, "删除成功")
}

// PageList 获取带分页的列表
func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1")) // strconv.Atoi() 转为 int 类型
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 前端渲染分页需要知道总条数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	reponse.Success(ctx, gin.H{"data": posts, "total": total}, "查询成功")
}
