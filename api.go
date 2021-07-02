package hcs

import (
	"github.com/gin-gonic/gin"
	"hcs/dao"
	"hcs/internal"
)

type PutAgentRequest struct {
	HostName string `json:"host_name"` // hostname
	Name     string `json:"name"`      // 主机名
	Version  string `json:"version"`   // agent版本
	Mark     string `json:"mark"`      // agent备注
}

type (
	HeartRequest struct {
		AgentID int `json:"agent_id"`
	}

	HeartResponse struct {
		Status int `json:"status"`
	}
)

type (
	GetTaskRequest struct {
		AgentID int `json:"agent_id"`
	}

	GetTaskResponse struct {
		Task []dao.Task `json:"task"`
	}
)

type (
	PutTaskRequest struct {
		AgentID int `json:"agent_id"`
		TaskID  int `json:"task_id"`
	}

	PutTaskResponse struct {
	}
)

func PutAgent(c *gin.Context) {
	p := new(PutAgentRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.SetResult(c, internal.ErrorCode, nil, err.Error())
		return
	}

	internal.SetResult(c, internal.SuccessCode, nil, internal.SuccessMsg)
}

func Heart(c *gin.Context) {
	p := new(HeartRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.SetResult(c, internal.ErrorCode, nil, err.Error())
		return
	}

	internal.SetResult(c, internal.SuccessCode, nil, internal.SuccessMsg)
}

func GetTask(c *gin.Context) {
	p := new(GetTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.SetResult(c, internal.ErrorCode, nil, err.Error())
		return
	}

	internal.SetResult(c, internal.SuccessCode, nil, internal.SuccessMsg)
}

func PutTask(c *gin.Context) {
	p := new(PutTaskRequest)
	if err := c.ShouldBindJSON(&p); err != nil {
		internal.SetResult(c, internal.ErrorCode, nil, err.Error())
		return
	}

	internal.SetResult(c, internal.SuccessCode, nil, internal.SuccessMsg)
}
