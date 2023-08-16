package router

import (
	"base_k8s/cloud/controller/api/v1/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// StartServer 启动 web server
func StartServer(r *gin.Engine) {
	// 路由分组：版本一
	version1 := r.Group("/api/v1")

	// api V1 的路由
	version1.GET("/hello", user.Show)
	// 启动 web server
	addr := fmt.Sprintf(":%v", viper.Get("app.port"))
	r.Run(addr)
}
