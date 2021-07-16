package manager

import (
	"hcs/dao"
)

type (
	GetTaskRequest struct {
		AgentID int `form:"agent_id"`
	}

	GetTaskResponse struct {
		Task []dao.TaskSimple `json:"task"`
	}
)

type (
	PutTaskRequest struct {
		Task []dao.AgentTaskSimple `json:"task"`
	}
)

type AgentTaskManager struct {
	agenttask *dao.AgentTaskDao
	task      *dao.TaskDao
}

func NewAgentTaskManager(agenttask *dao.AgentTaskDao, task *dao.TaskDao) *AgentTaskManager {
	at := new(AgentTaskManager)
	at.agenttask = agenttask
	at.task = task
	return at
}

func (at *AgentTaskManager) GetTask(agentId int) ([]dao.Task, error) {
	t, err := at.agenttask.FindTask(agentId)
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
	tasks, err := at.task.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// 给Agent增加任务
func (a *AgentTaskManager) AddAgentTask(agentId int, taskId int, mark string) (int, error) {
	t := dao.AgentTask{
		AgentId: agentId,
		TaskId:  taskId,
		Mark:    mark,
	}

	return a.agenttask.Insert(&t)
}

func (a *AgentTaskManager) SetStatus(task []dao.AgentTaskSimple) error {
	if len(task) == 0 {
		return nil
	}

	return a.agenttask.UpdateStatus(task)
}

func (a *AgentTaskManager) GetAll() ([]dao.AgentTask, error) {
	return a.agenttask.FindAll()
}

func (a *AgentTaskManager) Delete(relationId int) error {
	return a.agenttask.Delete(relationId)
}
