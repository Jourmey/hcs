package dao

import "time"

const DBTask = "task"

type Task struct {
	ID         int       `gorm:"column:id"`       // 主键ID
	Content    string    `gorm:"column:content"`  // 任务内容
	Status     int       `gorm:"column:status"`   // 任务状态
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"` // 记录更新时间
}

func (Task) TableName() string {
	return DBTask
}

// Status
const (
	Normal   int = iota //任务创建
	Running             // 任务运行中。
	Finished            // 任务执行成功。
	Failed              // 任务执行失败。
)
