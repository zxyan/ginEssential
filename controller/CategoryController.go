package controller

import (
	"ctjsoft/ginessential/model"
	"ctjsoft/ginessential/reponse"
	"ctjsoft/ginessential/repository"
	"ctjsoft/ginessential/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	// DB *gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	//db := common.GetDB()             // 获取数据库连接
	//db.AutoMigrate(model.Category{}) // 添加自动迁移
	//
	//// 创建新的 CategoryController 实例
	//return CategoryController{DB: db}

	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(&model.Category{})
	return CategoryController{Repository: repository}
}

// Create 新增
func (c CategoryController) Create(ctx *gin.Context) {
	// 使用结构体来解析请求参数
	// var requestCategory model.Category
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	//ctx.Bind(&requestCategory) // gin 框架提供的 Bind 函数
	//if requestCategory.Name == "" {
	//	reponse.Fail(ctx, nil, "数据验证错误, 分类名称必填")
	//	return
	//}
	// c.DB.Create(&requestCategory)

	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		reponse.Fail(ctx, nil, "数据验证错误, 分类名称必填")
		return
	}
	// category := model.Category{Name: requestCategory.Name}
	// c.DB.Create(&category)
	// reponse.Success(ctx, gin.H{"category": requestCategory}, "创建成功")

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}
	reponse.Success(ctx, gin.H{"category": category}, "创建成功")
}

// Update 修改
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定 body 种的参数
	// 使用结构体来解析请求参数
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory) // gin 框架提供的 Bind 函数
	//if requestCategory.Name == "" {
	//	reponse.Fail(ctx, nil, "数据验证错误, 分类名称必填")
	//	return
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		reponse.Fail(ctx, nil, "数据验证错误, 分类名称必填")
		return
	}

	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//var updateCategory model.Category // 定义变量, 需要更新的分类
	//if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
	//	reponse.Fail(ctx, nil, "分类不存在")
	//	return
	//}

	// 更新分类, 3 个参数: map, struct, name value
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	//reponse.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
	fmt.Println("categoryId:", categoryId)
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		reponse.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类
	category, crr := c.Repository.Update(*updateCategory, requestCategory.Name)
	if crr != nil {
		panic(crr)
		return
	}
	reponse.Success(ctx, gin.H{"category": category}, "修改成功")
}

// Show 获取
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//var category model.Category // 定义变量, 查询的分类
	//if c.DB.First(&category, categoryId).RecordNotFound() {
	//	reponse.Fail(ctx, nil, "分类不存在")
	//	return
	//}

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		reponse.Fail(ctx, nil, "分类不存在")
		return
	}

	reponse.Success(ctx, gin.H{"category": category}, "获取成功")
}

// Delete 删除
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
	//	reponse.Fail(ctx, nil, "删除失败, 请重试")
	//	return
	//}

	if err := c.Repository.DeleteById(categoryId); err != nil {
		reponse.Fail(ctx, nil, "删除失败, 请重试")
		return
	}

	reponse.Success(ctx, nil, "删除成功")
}
