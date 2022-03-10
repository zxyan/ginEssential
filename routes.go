package main

import (
	"ctjsoft/ginessential/controller"
	"ctjsoft/ginessential/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // 将中间件用来保护用户信息接口

	categoryRoutes := r.Group("/categories")                 // 创建路由分组
	categoryController := controller.NewCategoryController() // 创建 CategoryController
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	postRoutes := r.Group("/posts")                  // 创建路由分组
	postRoutes.Use(middleware.AuthMiddleware())      // 加上登录用户中间件
	postController := controller.NewPostController() // 创建 PostController
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.POST("page/list", postController.PageList)

	return r
}
