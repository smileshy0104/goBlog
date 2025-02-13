package response

import (
	"goBlog/lib/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 数据返回通用JSON数据结构
type Response struct {
	Code    int         `json:"code"` // 错误码((0:成功, -1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

func tpl(c *gin.Context, code int, msg string, data interface{}, t string) {
	responseData := interface{}(nil)
	if data != nil {
		responseData = data
	}
	switch t {
	case "xml":
		c.XML(http.StatusOK, Response{
			Code:    code,
			Message: msg,
			Data:    responseData,
		})
	case "yaml":
		c.YAML(http.StatusOK, Response{
			Code:    code,
			Message: msg,
			Data:    responseData,
		})
	case "toml":
		c.TOML(http.StatusOK, Response{
			Code:    code,
			Message: msg,
			Data:    responseData,
		})
	case "jsonp":
		c.JSONP(http.StatusOK, Response{
			Code:    code,
			Message: msg,
			Data:    responseData,
		})
	default:
		c.JSON(http.StatusOK, Response{
			Code:    code,
			Message: msg,
			Data:    responseData,
		})
	}
}

func Output(c *gin.Context, code int, msg string, data interface{}) {
	tpl(c, code, msg, data, global.GVA_OUTPUT)
}
