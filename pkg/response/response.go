package response

import (
	"github.com/gin-gonic/gin"
)

func Json(c *gin.Context, maps map[string]interface{}) {
	var data interface{}
	code := SUCCESS     //200
	msg := GetMsg(code) //ok
	state := true

	if item, ok := maps["state"].(bool); ok {
		state = item
	}

	if rows, ok := maps["data"]; ok {
		data = rows
	}

	if item, ok := maps["code"].(int); ok {
		code = item
		msg = GetMsg(code)
	}

	c.JSON(SUCCESS, gin.H{
		"state": state,
		"code":  code,
		"msg":   msg,
		"data":  data,
	})
}
