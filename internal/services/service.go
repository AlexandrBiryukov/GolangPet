package services

import (
	"errors"
)

type TaskService interface {
	CreateTask(taskText string) (Tasks, error)
	GetTasks() ([]Tasks, error)
	GetTaskById(id string) (Tasks, error)
	UpdateTask(id, task string) (Tasks, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepositoty
}

func NewTaskService(r TaskRepositoty) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(taskText string) (Tasks, error) {

	if taskText == "" {
		return Tasks{}, errors.New("Ошибка: поле task пустое или не обнаружено")
	}

	task := Tasks{
		//ID:   uuid.New(),
		Task: taskText,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Tasks{}, err
	}

	return task, nil
}

func (s *taskService) GetTasks() ([]Tasks, error) {
	return s.repo.GetTasks()
}

func (s *taskService) GetTaskById(id string) (Tasks, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id, taskText string) (Tasks, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Tasks{}, err
	}
	if taskText == "" {
		return Tasks{}, errors.New("Ошибка: поле task пустое или не обнаружено")
	}
	task.Task = taskText

	if err := s.repo.UpdateTask(task); err != nil {
		return Tasks{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
