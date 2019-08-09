package api

import (
	"github.com/gin-gonic/gin"
)

func configTestsRouter(router *gin.RouterGroup) {
	router.GET("/api/test", test)
}

func test(c *gin.Context)  {
	c.JSON(200,map[string]interface{}{
		"name":"framework",
		"pkg" : "golang gin gorm",
	})
}

//func GetAllBooks(c *gin.Context) {
//	maps := make(map[string]interface{})
//	data := make(map[string]interface{})
//
//	var total int
//	total = models.GetBookTotal(maps)
//	data["total"] = total
//
//	maps = request.GetPage(c, maps, false)
//	data["list"] = models.GetBooks(maps)
//
//	response.Json(c, response.JsonType{Data: data})
//}
//
//func GetBook(c *gin.Context) {
//	id := request.GetParam(c, "id", "")
//	book, _ := models.GetBook(id.MustInt())
//
//	if book.ID > 0 {
//		response.Json(c, response.JsonType{Data: book})
//	} else {
//		response.Json(c, response.JsonType{
//			State: false,
//			Code:  http.StatusNotFound,
//		})
//	}
//}
//
//func AddBook(c *gin.Context) {
//	//maps := make(map[string]interface{})
//
//	book := models.Book{}
//	json := response.JsonType{}
//
//	if err := models.AddBook(&book); err != nil {
//		json.Code = http.StatusInternalServerError
//		json.Msg = "新增失败"
//		json.State = false
//		json.Data = book
//	}
//	response.Json(c, json)
//}
//
//func UpdateBook(c *gin.Context) {
//	id := request.GetParam(c, "id", "")
//	maps := make(map[string]interface{})
//
//	book, _ := models.GetBook(id.MustInt())
//
//	json := response.JsonType{}
//	if book.ID > 0 {
//		if err := models.UpdateBook(&book, maps); err != nil {
//			json.Code = http.StatusInternalServerError
//			json.Msg = "修改失败"
//			json.State = false
//		}
//	} else {
//		json.Code = http.StatusNotFound
//		json.State = false
//	}
//
//	response.Json(c, json)
//}
//
//func DeleteBook(c *gin.Context) {
//	id := request.GetParam(c, "id", "")
//	json := response.JsonType{}
//	if err := models.DeleteBook(id.MustInt()); err != nil {
//		json.Code = http.StatusInternalServerError
//		json.Msg = "删除失败"
//		json.State = false
//	}
//	response.Json(c, json)
//}
