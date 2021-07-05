package agenttask

import (
	"github.com/jinzhu/gorm"
	"hcs/dao"
	"time"
)

const DBName = "agent_task"

type AgentTask struct {
	ID         int       `gorm:"column:id"`       // 主键ID
	AgentId    int       `gorm:"column:agent_id"` // hostname
	TaskId     int       `gorm:"column:task_id"`
	Status     int       `gorm:"column:status"`      // 任务状态
	Mark       string    `gorm:"column:mark"`        // 关系备注
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
	ID     int `json:"id"`     // 主键ID
	Status int `json:"status"` // 任务状态
}

func Insert(t *AgentTask) (int, error) {
	err := dao.DB().Create(t).Error
	return t.ID, err
}

func find(fu func(d *gorm.DB) *gorm.DB) (t *AgentTask, err error) {
	err = fu(dao.DB()).
		Where("delete_flag = ?", 0).
		First(t).Error
	return
}

func finds(fu func(d *gorm.DB) *gorm.DB) (t []AgentTask, err error) {
	err = fu(dao.DB()).
		Where("delete_flag = ?", 0).
		Find(&t).Error
	return
}

func FindTask(agentId int) ([]AgentTask, error) {
	return finds(func(d *gorm.DB) *gorm.DB {
		return d.Where("agent_id = ?", agentId).
			Where("status != ?", Finished)
	})
}

func FindStatusMin(agentId int) (*AgentTask, error) {
	return find(func(d *gorm.DB) *gorm.DB {
		return d.Where("agent_id = ?", agentId).
			Where("status != ?", Finished).
			Order("status DESC")
	})
}

func UpdateStatus(tasks []AgentTaskSimple) error {
	m := make(map[int][]int, 0)
	for i := 0; i < len(tasks); i++ {
		m[tasks[i].Status] = append(m[tasks[i].Status], tasks[i].ID)
	}
	for status, ids := range m {
		err := dao.DB().Update("status = ?", status).Where("id in (?)", ids).Error
		if err != nil {
			return err
		}
	}
	return nil
}
