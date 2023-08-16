package user

import (
	"base_k8s/cloud/service/user"
	"github.com/gin-gonic/gin"
)

func Show(ctx *gin.Context) {
	user.Show(ctx)
}
