package manager

import (
	"hcs/dao"
)

type AgentManager struct {
	agent     *dao.AgentDao
	agenttask *dao.AgentTaskDao
}

func NewAgentManager(agent *dao.AgentDao, agenttask *dao.AgentTaskDao) *AgentManager {
	a := new(AgentManager)
	a.agenttask = agenttask
	a.agent = agent
	return a
}

func (a *AgentManager) PutAgent(hostName string, name string, version string, mark string) (int, error) {
	t := dao.Agent{
		HostName: hostName,
		Name:     name,
		Version:  version,
		Mark:     mark,
	}

	return a.agent.Insert(&t)
}

func (a *AgentManager) Heart(agentID int) (bool, error) {
	err := a.agent.UpdateHeart(agentID)
	if err != nil {
		return false, err
	}

	t, err := a.agenttask.Exist(agentID, dao.Normal)
	return t, nil
}

func (a *AgentManager) GetAllAgent() ([]dao.Agent, error) {
	return a.agent.GetAll()
}
