package response

import (
	"github.com/gin-gonic/gin"
)

type JsonType struct {
	Code  int
	Msg   string
	State bool
	Data  interface{}
}

func Json(c *gin.Context, jsonType JsonType) {
	var data interface{}
	code := SUCCESS     //200
	msg := GetMsg(code) //ok
	state := true

	if jsonType.State == false {
		state = false
	}

	if jsonType.Data != nil {
		data = jsonType.Data
	}

	if jsonType.Code != 0 {
		code = jsonType.Code
		msg = GetMsg(code)
	}

	if jsonType.Msg != "" {
		msg = jsonType.Msg
	}

	c.JSON(SUCCESS, gin.H{
		"state": state,
		"code":  code,
		"msg":   msg,
		"data":  data,
	})
}
