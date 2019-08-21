package request

import (
	"framework/config"
	"framework/pkg/utli"
	"github.com/gin-gonic/gin"
	"github.com/openset/php2go/php"
)

func GetParam(c *gin.Context, key, ext string) utli.StrTo {
	v := c.Param(key + ext)
	v = php.StrReplace(ext, "", v, 1)
	return utli.StrTo(v)
}

func GetPage(c *gin.Context, maps map[string]interface{}, must bool) map[string]interface{} {

	if pageS, ok := c.GetQuery("page"); ok || must {
		page := utli.StrTo(pageS).MustInt()
		size := utli.StrTo(c.Query("size")).MustInt()
		if page < 1 {
			page = 1
		}
		maxSize := config.GetApp("MAX_PAGE_SIZE").MustInt()
		if size < 1 {
			size = config.PageSize
		} else if size > maxSize {
			size = maxSize
		}
		maps["page"] = (page - 1) * size
		maps["size"] = size
	}

	return maps
}
