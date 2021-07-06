package hcs

import (
	"github.com/gin-gonic/gin"
	"hcs/dao"
)

func WrapGroup(router *gin.RouterGroup) {
	router.PUT("/agent", PutAgent)    // agent注册
	router.PUT("/agent/heart", Heart) // agent心跳
	router.GET("/task", GetTask)      // agent获取任务
	router.PUT("/task", PutTask)      // agent回写任务
}

type MySql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Hostport string `json:"hostport"`
	Database string `json:"database"`
}

func MustInitDB(db MySql) {
	dao.MustInitDB(db.Username, db.Password, db.Hostname, db.Hostport, db.Database)
}

