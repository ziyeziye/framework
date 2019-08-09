package api

import (
	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.RouterGroup) {
	configTestsRouter(router)
}
