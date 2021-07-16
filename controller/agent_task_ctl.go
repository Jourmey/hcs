package controller

import (
	"github.com/gin-gonic/gin"
	"hcs/dao"
	"hcs/manager"
	"strconv"
)

type AgentTaskCtl struct {
	agentTaskManager *manager.AgentTaskManager
}

func NewAgentTaskCtl(agentTaskManager *manager.AgentTaskManager) *AgentTaskCtl {
	a := new(AgentTaskCtl)
	a.agentTaskManager = agentTaskManager
	return a
}

func (a *AgentTaskCtl) AddRelation(c *gin.Context) {
	type param struct {
		AgentId int    `json:"agent_id"`
		TaskId  int    `json:"task_id"`
		Mark    string `json:"mark"`
	}
	var p param
	if err := c.ShouldBindJSON(&p); err != nil {
		Return(c, nil, err)
		return
	}

	id, err := a.agentTaskManager.AddAgentTask(p.AgentId, p.TaskId, p.Mark)
	Return(c, id, err)
}
func (a *AgentTaskCtl) DeleteRelation(c *gin.Context) {
	relation_id := c.Query("relation_id")
	relationId, err := strconv.Atoi(relation_id)
	if err != nil {
		Return(c, nil, err)
		return
	}
	err = a.agentTaskManager.Delete(relationId)
	Return(c, nil, err)
}

func (a *AgentTaskCtl) GetAll(c *gin.Context) {
	res, err := a.agentTaskManager.GetAll()
	Return(c, res, err)
}

// agent回写任务状态
func (a *AgentTaskCtl) PostStatus(c *gin.Context) {
	p := new(manager.PutTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		Return(c, nil, err)
		return
	}

	err := a.agentTaskManager.SetStatus(p.Task)
	Return(c, "", err)
}

func (a *AgentTaskCtl) GetRelationByAgent(c *gin.Context) {
	agent_id := c.Query("agent_id")
	agentId, err := strconv.Atoi(agent_id)
	if err != nil {
		Return(c, nil, err)
		return
	}
	res, err := a.agentTaskManager.GetTask(agentId)
	r := manager.GetTaskResponse{Task: make([]dao.TaskSimple, 0, len(res))}
	for i := 0; i < len(res); i++ {
		r.Task = append(r.Task, dao.TaskSimple{
			ID:      res[i].ID,
			Content: res[i].Content,
		})
	}

	Return(c, r, err)
}
func (a *AgentTaskCtl) GetRelationByTask(c *gin.Context) {

}
