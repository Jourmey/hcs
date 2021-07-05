package hcs

import (
	"github.com/gin-gonic/gin"
	"hcs/dao"
	agenttask "hcs/dao/agent_task"
	"hcs/dao/task"
	"hcs/internal"
)

func WrapGroup(router *gin.RouterGroup) {
	router.PUT("/hcs/agent", PutAgent)    // agent注册
	router.PUT("/hcs/agent/heart", Heart) // agent心跳
	router.GET("/hcs/task", GetTask)      // agent获取任务
	router.PUT("/hcs/task", PutTask)      // agent回写任务
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

type Task struct {
	Content string // 任务内容
}

type AgentTask struct {
	AgentId int    `gorm:"column:agent_id"` // hostname
	TaskId  int    `gorm:"column:task_id"`
	Mark    string `gorm:"column:mark"` // 关系备注
}

// 新建任务
func AddTask(t *Task) (int, error) {
	tt := task.Task{
		Content: t.Content,
	}
	return internal.AddTask(&tt)
}

// 给Agent增加任务
func AddAgentTask(t *AgentTask) (int, error) {
	tt := agenttask.AgentTask{
		AgentId: t.AgentId,
		TaskId:  t.TaskId,
		Mark:    t.Mark,
	}

	return internal.AddAgentTask(&tt)
}
