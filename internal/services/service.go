package services

import (
	"errors"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(taskText string) (RequestBody, error)
	GetTasks() ([]RequestBody, error)
	GetTaskById(id string) (RequestBody, error)
	UpdateTask(id, task string) (RequestBody, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepositoty
}

func NewTaskService(r TaskRepositoty) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(taskText string) (RequestBody, error) {

	if taskText == "" {
		return RequestBody{}, errors.New("Ошибка: поле task пустое или не обнаружено")
	}

	task := RequestBody{
		ID:   uuid.NewString(),
		Task: taskText,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return RequestBody{}, err
	}

	return task, nil
}

func (s *taskService) GetTasks() ([]RequestBody, error) {
	return s.repo.GetTasks()
}

func (s *taskService) GetTaskById(id string) (RequestBody, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id, taskText string) (RequestBody, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return RequestBody{}, err
	}
	if taskText == "" {
		return RequestBody{}, errors.New("Ошибка: поле task пустое или не обнаружено")
	}
	task.Task = taskText

	if err := s.repo.UpdateTask(task); err != nil {
		return RequestBody{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
