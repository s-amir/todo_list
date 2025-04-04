package task

import "awesomeProject/server/entity"

type TaskRepository interface {
	CreateNewTask(t entity.Task) (*entity.Task, error)
	ListUserTasks(userId int) (*[]entity.Task, error)
}

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(tr TaskRepository) *TaskService {
	return &TaskService{repository: tr}
}

type CreateRequest struct {
	Title               string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task entity.Task
}

func (t *TaskService) Create(requestTask CreateRequest) (*CreateResponse, error) {
	newTask, err := t.repository.CreateNewTask(
		entity.Task{
			Title:      requestTask.Title,
			DueDate:    requestTask.DueDate,
			CategoryID: requestTask.CategoryID,
			IsDone:     false,
			UserID:     requestTask.AuthenticatedUserID})
	if err != nil {
		return nil, err
	}
	return &CreateResponse{Task: *newTask}, nil
}

type ListUserTasksRequest struct {
	UserID int
}

type ListUserTasksResponse struct {
	Tasks []entity.Task
}

func (t *TaskService) List(request ListUserTasksRequest) (*ListUserTasksResponse, error) {
	listUserTasks, err := t.repository.ListUserTasks(request.UserID)
	if err != nil {
		return nil, err
	}
	return &ListUserTasksResponse{Tasks: *listUserTasks}, nil
}
