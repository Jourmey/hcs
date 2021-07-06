package agenttask

import (
	"github.com/jinzhu/gorm"
	"hcs/dao"
	"time"
)

const DBName = "agent_task"

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
	return DBName
}

// Status
const (
	Normal   int = iota //任务创建
	Running             // 任务运行中。
	Failed              // 任务执行失败。
	Finished            // 任务执行成功。
)

type AgentTaskSimple struct {
	ID     int    `json:"id"`     // 主键ID
	Status int    `json:"status"` // 任务状态
	Reason string `json:"reason"` // 任务状态原因
}

func Insert(t *AgentTask) (int, error) {
	err := dao.DB().Create(t).Error
	return t.ID, err
}

func find(fu func(d *gorm.DB) *gorm.DB) (*AgentTask, error) {
	t := new(AgentTask)
	err := fu(dao.DB()).
		Where("delete_flag = ?", 0).
		First(t).Error
	return t, err
}

func finds(fu func(d *gorm.DB) *gorm.DB) ([]AgentTask, error) {
	var t []AgentTask
	err := fu(dao.DB()).
		Where("delete_flag = ?", 0).
		Find(&t).Error
	return t, err
}

func FindTask(agentId int) ([]AgentTask, error) {
	return finds(func(d *gorm.DB) *gorm.DB {
		return d.Where("agent_id = ?", agentId).
			Where("status = ?", Normal)
	})
}

func Exist(agentId int, status int) (bool, error) {
	r, err := find(func(d *gorm.DB) *gorm.DB {
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

func UpdateStatus(tasks []AgentTaskSimple) error {
	for i := 0; i < len(tasks); i++ {
		err := dao.DB().Table(DBName).
			Update("status", tasks[i].Status).
			Update("reason", tasks[i].Reason).
			Where("id = ?", tasks[i].ID).Error
		if err != nil {
			return err
		}
	}

	return nil
}
