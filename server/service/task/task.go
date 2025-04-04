package service

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

type createRequest struct {
	Title               string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	task entity.Task
}

func (t *TaskService) Create(requestTask createRequest) (*CreateResponse, error) {
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
	return &CreateResponse{task: *newTask}, nil
}

type ListUserTasksRequest struct {
	UserID int
}

type ListUserTasksResponse struct {
	Tasks []entity.Task
}

func (t *TaskService) ListUserTasks(request ListUserTasksRequest) (*ListUserTasksResponse, error) {
	listUserTasks, err := t.repository.ListUserTasks(request.UserID)
	if err != nil {
		return nil, err
	}
	return &ListUserTasksResponse{Tasks: *listUserTasks}, nil
}
