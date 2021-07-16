package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

const DBAgentTaskName = "agent_task"

type AgentTask struct {
	ID         int       `gorm:"column:id"` // 主键ID
	AgentId    int       `gorm:"column:agent_id"`
	TaskId     int       `gorm:"column:task_id"`
	Mark       string    `gorm:"column:mark"`        // 关系备注
	Status     int       `gorm:"column:status"`      // 任务状态 0:任务创建 1:任务运行中 2:任务执行失败 3:任务执行成功
	Reason     string    `gorm:"column:reason"`      // 任务状态原因
	DeleteFlag int       `gorm:"column:delete_flag"` // 删除状态，0 正常、1 删除
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"` // 记录更新时间
}

func (AgentTask) TableName() string {
	return DBAgentTaskName
}

// Status
const (
	Normal   int = iota //任务创建
	Running             // 任务运行中。
	Failed              // 任务执行失败。
	Finished            // 任务执行成功。
)

type AgentTaskDao struct{}

type AgentTaskSimple struct {
	ID     int    `json:"id"`     // 主键ID
	Status int    `json:"status"` // 任务状态
	Reason string `json:"reason"` // 任务状态原因
}

func (a *AgentTaskDao) Insert(t *AgentTask) (int, error) {
	err := DB().Create(t).Error
	return t.ID, err
}

func (a *AgentTaskDao) find(fu func(d *gorm.DB) *gorm.DB) (*AgentTask, error) {
	t := new(AgentTask)
	err := fu(DB()).
		Where("delete_flag = ?", 0).
		First(t).Error
	return t, err
}

func (a *AgentTaskDao) finds(fu func(d *gorm.DB) *gorm.DB) ([]AgentTask, error) {
	var t []AgentTask
	err := fu(DB()).
		Where("delete_flag = ?", 0).
		Find(&t).Error
	return t, err
}

func (a *AgentTaskDao) FindTask(agentId int) ([]AgentTask, error) {
	return a.finds(func(d *gorm.DB) *gorm.DB {
		return d.Where("agent_id = ?", agentId).
			Where("status = ?", Normal)
	})
}

func (a *AgentTaskDao) Exist(agentId int, status int) (bool, error) {
	r, err := a.find(func(d *gorm.DB) *gorm.DB {
		return d.Where("agent_id = ?", agentId).
			Where("status = ?", status)
	})
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return r.ID != 0, nil
}

func (a *AgentTaskDao) UpdateStatus(tasks []AgentTaskSimple) error {
	for i := 0; i < len(tasks); i++ {
		err := DB().Table(DBAgentTaskName).
			Update("status", tasks[i].Status).
			Update("reason", tasks[i].Reason).
			Where("id = ?", tasks[i].ID).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *AgentTaskDao) FindAll() ([]AgentTask, error) {
	return a.finds(func(d *gorm.DB) *gorm.DB {
		return d
	})
}

func (a *AgentTaskDao) Delete(relationId int) error {
	err := DB().Table(DBAgentTaskName).
		Update("delete_flag", 1).
		Where("id = ?", relationId).Error
	return err
}
