package main

import (
	"ctjsoft/ginessential/common"
	"github.com/spf13/viper"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitConfig() // 读取配置文件

	db := common.InitDB()
	defer db.Close()

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了 request 和 response
	r = CollectRoute(r)
	// 3.监听端口，默认在 8080
	// Run("里面不指定端口号默认为 8080")
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(":8000"))
}

// InitConfig 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()                                  // 获取当前的工作目录
	viper.SetConfigName("application")                        // 设置读取的文件名
	viper.SetConfigType("yml")                                // 设置读取的文件的类型
	viper.AddConfigPath(workDir + "/src/ginEssential/config") // 设置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
