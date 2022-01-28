package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 处理跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8083") // 允许访问的域名
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")                      // 设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")                    // 设置可以通过访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")                    // 允许请求带的 header 信息
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200) // 判断是 options 请求, 直接返回 200
		} else {
			ctx.Next() // 中间件向下传递
		}
	}
}
