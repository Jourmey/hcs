package internal

import (
	agenttask "hcs/dao/agent_task"
	"hcs/dao/task"
)

// 新建任务
func AddTask(t *task.Task) (int, error) {
	return task.Insert(t)
}

// 给Agent增加任务
func AddAgentTask(t *agenttask.AgentTask) (int, error) {
	return agenttask.Insert(t)
}
