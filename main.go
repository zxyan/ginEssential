package main

import (
	"ctjsoft/ginessential/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了 request 和 response
	r = CollectRoute(r)
	// 3.监听端口，默认在 8080
	// Run("里面不指定端口号默认为 8080")
	panic(r.Run(":8000"))
}
