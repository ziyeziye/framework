package fm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestTpl(c *gin.Context) {
	c.HTML(http.StatusOK, "test.tpl", gin.H{
		"title":    "framework test",
		"name":    "golang gin gorm",
	})

}
