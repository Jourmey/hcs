package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

const DBAgentName = "agent"

type Agent struct {
	ID         int       `gorm:"column:id"`         // 主键ID
	HostName   string    `gorm:"column:host_name"`  // hostname
	Name       string    `gorm:"column:name"`       // 主机名
	Version    string    `gorm:"column:version"`    // agent版本
	Mark       string    `gorm:"column:mark"`       // agent备注
	HeartTime  time.Time `gorm:"column:heart_time"` // 心跳时间
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"` // 记录更新时间
}

func (Agent) TableName() string {
	return DBAgentName
}

type AgentDao struct{}

func (a *AgentDao) find(fu func(d *gorm.DB) *gorm.DB) (t *Agent, err error) {
	err = fu(DB()).First(t).Error
	return
}

func (a *AgentDao) finds(fu func(d *gorm.DB) *gorm.DB) (t []Agent, err error) {
	err = fu(DB()).Find(&t).Error
	return
}

func (a *AgentDao) Insert(agent *Agent) (int, error) {
	err := DB().Omit("create_time", "update_time", "heart_time").Create(agent).Error
	return agent.ID, err
}

func (a *AgentDao) UpdateHeart(id int) error {
	err := DB().
		Table(DBAgentName).
		Update("heart_time", time.Now()).
		Where("id = ?", id).Error
	return err
}

func (a *AgentDao) GetAll() ([]Agent, error) {
	return a.finds(func(d *gorm.DB) *gorm.DB {
		return d
	})
}
