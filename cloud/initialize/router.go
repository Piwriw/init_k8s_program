package initialize

import "github.com/gin-gonic/gin"

// Router 初始化 gin router
func Router() *gin.Engine {
	return gin.Default()
}
