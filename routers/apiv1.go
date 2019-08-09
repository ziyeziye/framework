package routers

import (
	"github.com/gin-gonic/gin"
)

func apiv1(r *gin.Engine) {
	group := r.Group("")
	{
		group.GET("/test", func(c *gin.Context) {
			c.JSON(200,map[string]interface{}{
				"name":"framework",
				"pkg" : "golang gin gorm",
			})
		})
		//api.ConfigRouter(group)
	}
	//apiv1 := r.Group("/api/v1")
	//{
	//获取标签列表
	//apiv1.GET("/tags", v1.GetTags)
	//新建标签
	//apiv1.POST("/tags", v1.AddTag)
	//更新指定标签
	//apiv1.PUT("/tags/:id", v1.EditTag)
	//删除指定标签
	//apiv1.DELETE("/tags/:id", v1.DeleteTag)
	//}

}
