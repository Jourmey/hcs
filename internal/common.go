package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func SetResult(c *gin.Context, code int, data interface{}, msg string) {
	res := new(Result)
	res.Code = code
	res.Result = data
	res.Msg = msg
	c.JSON(http.StatusOK, res)
}
