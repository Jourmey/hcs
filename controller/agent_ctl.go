package controller

import (
	"github.com/gin-gonic/gin"
	"hcs/manager"
)

type AgentCtl struct {
	agentManager *manager.AgentManager
}

func NewAgentCtl(agentManager *manager.AgentManager) *AgentCtl {
	a := new(AgentCtl)
	a.agentManager = agentManager
	return a
}

// POST agent注册
func (a *AgentCtl) AddAgent(c *gin.Context) {
	p := new(PostAgentRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		Return(c, nil, err)
		return
	}

	res, err := a.agentManager.PutAgent(p.HostName, p.Name, p.Version, p.Mark)
	Return(c, PostAgentResponse{AgentID: res}, err)
}

// Get agent心跳
func (a *AgentCtl) Heart(c *gin.Context) {
	p := new(HeartRequest)
	if err := c.ShouldBindQuery(&p); err != nil {
		Return(c, nil, err)
		return
	}

	status, err := a.agentManager.Heart(p.AgentID)
	Return(c, HeartResponse{Status: status}, err)
}

// GET 获取agent列表
func (a *AgentCtl) GetAll(c *gin.Context) {
	agents, err := a.agentManager.GetAllAgent()
	Return(c, agents, err)
}

type (
	PostAgentRequest struct {
		HostName string `json:"host_name"` // hostname
		Name     string `json:"name"`      // 主机名
		Version  string `json:"version"`   // agent版本
		Mark     string `json:"mark"`      // agent备注
	}
	PostAgentResponse struct {
		AgentID int `json:"agent_id"`
	}
)

type (
	HeartRequest struct {
		AgentID int `form:"agent_id"`
	}

	HeartResponse struct {
		Status bool `json:"status"`
	}
)
