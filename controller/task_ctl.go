package controller

import (
	"github.com/gin-gonic/gin"
	"hcs/manager"
	"strconv"
)

type TaskCtl struct {
	taskManager *manager.TaskManager
}

func NewTaskCtl(taskManager *manager.TaskManager) *TaskCtl {
	t := new(TaskCtl)
	t.taskManager = taskManager
	return t
}

// 新建任务
func (t *TaskCtl) AddTask(c *gin.Context) {
	p := new(PostTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		Return(c, nil, err)
		return
	}

	res, err := t.taskManager.AddTask(p.Content)
	Return(c, PostAgentResponse{AgentID: res}, err)
}

// 获取任务列表
func (t *TaskCtl) GetAll(c *gin.Context) {
	all, err := t.taskManager.GetAll()
	Return(c, all, err)
}

// 删除任务
func (t *TaskCtl) DeleteTask(c *gin.Context) {
	taskId := c.Query("task_id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		Return(c, nil, err)
		return
	}
	err = t.taskManager.Delete(id)
	Return(c, nil, err)
}

type PostTaskRequest struct {
	Content string `json:"content"` // 任务内容
}
