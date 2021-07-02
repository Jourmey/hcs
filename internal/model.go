package internal

import (
	"hcs/dao/task"
)



type (
	PutAgentRequest struct {
		HostName string `json:"host_name"` // hostname
		Name     string `json:"name"`      // 主机名
		Version  string `json:"version"`   // agent版本
		Mark     string `json:"mark"`      // agent备注
	}
	PutAgentResponse struct {
		AgentID int `json:"agent_id"`
	}
)

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
		Task []task.Task `json:"task"`
	}
)

type (
	PutTaskRequest struct {
		Task []PutTaskInfo
	}
	PutTaskInfo struct {
		TaskID int `json:"task_id"` // 主键ID
		Status int `json:"status"`  // 任务状态
	}

	//PutTaskResponse struct {
	//}
)