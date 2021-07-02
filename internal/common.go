package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func Return(c *gin.Context, data interface{}, err error) {
	res := new(Result)
	if err == nil {
		res.Code = 0
		res.Result = data
		res.Msg = "success"
	} else {
		res.Code = 1
		res.Result = data
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res)
}
