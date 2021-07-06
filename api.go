package hcs

import (
	"github.com/gin-gonic/gin"
	agenttask "hcs/dao/agent_task"
	"hcs/dao/task"
	"hcs/internal"
)

func PutAgent(c *gin.Context) {
	p := new(internal.PutAgentRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	res, err := internal.PutAgent(p)
	internal.Return(c, internal.PutAgentResponse{AgentID: res}, err)
}

func Heart(c *gin.Context) {
	p := new(internal.HeartRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	status, err := internal.Heart(p)
	internal.Return(c, internal.HeartResponse{Status: status}, err)
}

func GetTask(c *gin.Context) {
	p := new(internal.GetTaskRequest)
	if err := c.ShouldBindQuery(p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	res, err := internal.GetTask(p)
	r := internal.GetTaskResponse{Task: make([]task.TaskSimple, 0, len(res))}
	for i := 0; i < len(res); i++ {
		r.Task = append(r.Task, task.TaskSimple{
			ID:      res[i].ID,
			Content: res[i].Content,
		})
	}

	internal.Return(c, r, err)
}

func PutTask(c *gin.Context) {
	p := new(internal.PutTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.Return(c, nil, err)
		return
	}

	err := internal.PutTask(p)
	internal.Return(c, "", err)
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
