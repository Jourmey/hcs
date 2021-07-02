package internal

import (
	"hcs/dao/agent"
	"hcs/dao/task"
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

func Heart(p *HeartRequest) (int, error) {
	return task.Finished, nil

}
func GetTask(p *GetTaskRequest) ([]task.Task, error) {
	return nil, nil

}
func PutTask(p *PutTaskRequest) error {
	return nil

}
