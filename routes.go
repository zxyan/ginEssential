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
	CategoryController := controller.NewCategoryController() // 创建 CategoryController
	categoryRoutes.POST("", CategoryController.Create)
	categoryRoutes.PUT("/:id", CategoryController.Update)
	categoryRoutes.GET("/:id", CategoryController.Show)
	categoryRoutes.DELETE("/:id", CategoryController.Delete)

	return r
}
