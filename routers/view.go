package routers

import (
	fm "framework/controllers"
	"github.com/gin-gonic/gin"
)

func view(r *gin.Engine) {
	r.GET("/", fm.TestTpl)
}
