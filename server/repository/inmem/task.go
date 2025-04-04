package inmem

import "awesomeProject/server/entity"

type Task struct {
	tasks []entity.Task
}

func NewTask() *Task {
	return &Task{}
}

func (task *Task) CreateNewTask(t entity.Task) (*entity.Task, error) {
	t.ID = len(task.tasks) + 1
	task.tasks = append(task.tasks, t)
	return &t, nil

}

func (task *Task) ListUserTasks(userId int) (*[]entity.Task, error) {
	var list []entity.Task
	for _, task := range task.tasks {
		if task.UserID == userId {
			list = append(list, task)
		}
	}
	return &list, nil
}
