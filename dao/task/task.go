package task

import (
	"github.com/jinzhu/gorm"
	"hcs/dao"
	"time"
)

const DBName = "task"

type Task struct {
	ID         int       `gorm:"column:id"`          // 主键ID
	Content    string    `gorm:"column:content"`     // 任务内容
	DeleteFlag int       `gorm:"column:delete_flag"` // 删除状态，0 正常、1 删除
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"` // 记录更新时间
}

func (Task) TableName() string {
	return DBName
}

func Insert(t *Task) (int, error) {
	err := dao.DB().Create(t).Error
	return t.ID, err
}

func FindByIds(ids []int) ([]Task, error) {
	return finds(func(d *gorm.DB) *gorm.DB {
		return d.Where("id in (?)", ids)
	})
}

func find(fu func(d *gorm.DB) *gorm.DB) (t *Task, err error) {
	err = fu(dao.DB()).
		Where("delete_flag = ?", 0).
		First(t).Error
	return
}

func finds(fu func(d *gorm.DB) *gorm.DB) (t []Task, err error) {
	err = fu(dao.DB()).
		Where("delete_flag = ?", 0).
		Find(&t).Error
	return
}
