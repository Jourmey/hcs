package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

const DBTaskName = "task"

type Task struct {
	ID         int       `gorm:"column:id"`          // 主键ID
	Content    string    `gorm:"column:content"`     // 任务内容
	DeleteFlag int       `gorm:"column:delete_flag"` // 删除状态，0 正常、1 删除
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"` // 记录更新时间
}

func (Task) TableName() string {
	return DBTaskName
}

type TaskSimple struct {
	ID      int    `json:"id"`      // 主键ID
	Content string `json:"content"` // 任务内容
}

type TaskDao struct{}

func (td *TaskDao) Insert(t *Task) (int, error) {
	err := DB().Create(t).Error
	return t.ID, err
}

func (td *TaskDao) FindByIds(ids []int) ([]Task, error) {
	return td.finds(func(d *gorm.DB) *gorm.DB {
		return d.Where("id in (?)", ids)
	})
}

func (td *TaskDao) find(fu func(d *gorm.DB) *gorm.DB) (t *Task, err error) {
	err = fu(DB()).
		Where("delete_flag = ?", 0).
		First(t).Error
	return
}

func (td *TaskDao) finds(fu func(d *gorm.DB) *gorm.DB) (t []Task, err error) {
	err = fu(DB()).
		Where("delete_flag = ?", 0).
		Find(&t).Error
	return
}

// 逻辑删除
func (td *TaskDao) Delete(id int) error {
	err := DB().Table(DBAgentTaskName).
		Update("delete_flag", 1).
		Where("id = ?", id).Error
	return err
}

func (td *TaskDao) FindAll() ([]Task, error) {
	return td.finds(func(d *gorm.DB) *gorm.DB {
		return d
	})
}
