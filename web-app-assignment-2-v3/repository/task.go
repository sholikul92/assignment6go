package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	if task == nil {
		return fmt.Errorf("task cannot be nil")
	}

	taskID := task.ID

	err := t.filebased.UpdateTask(taskID, *task)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Delete(id int) error {
	err := t.filebased.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}


func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks, err := t.filebased.GetTasks()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	categories, err := t.filebased.GetTaskListByCategory(id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
