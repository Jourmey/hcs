package manager

import (
	"hcs/dao"
)

type TaskManager struct {
	task *dao.TaskDao
}

func NewTaskManager(task *dao.TaskDao) *TaskManager {
	a := new(TaskManager)
	a.task = task
	return a
}

// 新建任务
func (a *TaskManager) AddTask(content string) (int, error) {
	t := dao.Task{
		Content: content,
	}
	return a.task.Insert(&t)
}

func (a *TaskManager) Delete(id int) error {
	return a.task.Delete(id)
}

func (a *TaskManager) GetAll() ([]dao.Task, error) {
	return a.task.FindAll()
}
