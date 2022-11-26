package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}

type WithData struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func Ok(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"ok":   true,
	})
}

func OkWithData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"ok":   true,
		"data": data,
	})
}

func Fail(c *gin.Context, code, errCode int, msg string) {
	c.JSON(code, gin.H{
		"code": errCode,
		"msg":  msg,
		"ok":   false,
	})
}

func InternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  "internal err",
		"ok":   false,
	})
}
