package dao

import "time"

const DBAgent = "agent"

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
	return DBAgent
}
