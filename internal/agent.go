package internal

import (
	"hcs/dao/agent"
	"hcs/dao/agent_task"
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
		Status bool `json:"status"`
	}
)

type (
	GetTaskRequest struct {
		AgentID int `form:"agent_id"`
	}

	GetTaskResponse struct {
		Task []task.TaskSimple `json:"task"`
	}
)

type (
	PutTaskRequest struct {
		Task []agenttask.AgentTaskSimple `json:"task"`
	}

	//PutTaskResponse struct {
	//}
)

func PutAgent(p *PutAgentRequest) (int, error) {
	t := agent.Agent{
		HostName: p.HostName,
		Name:     p.Name,
		Version:  p.Version,
		Mark:     p.Mark,
	}

	return agent.Insert(&t)
}

func Heart(p *HeartRequest) (bool, error) {
	err := agent.UpdateHeart(p.AgentID)
	if err != nil {
		return false, err
	}

	t, err := agenttask.Exist(p.AgentID, agenttask.Normal)
	return t, nil
}

func GetTask(p *GetTaskRequest) ([]task.Task, error) {
	t, err := agenttask.FindTask(p.AgentID)
	if err != nil {
		return nil, err
	}
	if len(t) == 0 {
		return nil, nil
	}

	ids := make([]int, 0, len(t))
	for i := 0; i < len(t); i++ {
		ids = append(ids, t[i].TaskId)
	}
	tasks, err := task.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func PutTask(p *PutTaskRequest) error {
	if len(p.Task) == 0 {
		return nil
	}

	return agenttask.UpdateStatus(p.Task)
}
