package hcs

import (
	"github.com/gin-gonic/gin"
)

func WrapGroup(router *gin.RouterGroup) {
	router.PUT("/hcs/agent", PutAgent)    // agent注册
	router.PUT("/hcs/agent/heart", Heart) // agent心跳
	router.GET("/hcs/task", GetTask)      // agent获取任务
	router.PUT("/hcs/task", PutTask)      // agent回写任务
}
