package services

import "errors"

type TaskService interface {
	CreateTask(taskText string, isDone bool, userID uint) (Tasks, error)
	GetTasks() ([]Tasks, error)
	GetTaskById(id string) (Tasks, error)
	PatchTask(id string, taskText *string, isDone *bool) (Tasks, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID uint) ([]Tasks, error)
}

type taskService struct {
	repo TaskRepositoty
}

func NewTaskService(r TaskRepositoty) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(taskText string, isDone bool, userID uint) (Tasks, error) {
	if taskText == "" {
		return Tasks{}, errors.New("Ошибка: поле task пустое или не обнаружено")
	}

	task := Tasks{
		Task:   taskText,
		IsDone: isDone,
		UserID: userID,
	}

	if err := s.repo.CreateTask(&task); err != nil {
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

func (s *taskService) PatchTask(id string, taskText *string, isDone *bool) (Tasks, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Tasks{}, err
	}

	if taskText != nil {
		if *taskText == "" {
			return Tasks{}, errors.New("Ошибка: поле task пустое")
		}
		task.Task = *taskText
	}

	if isDone != nil {
		task.IsDone = *isDone
	}

	if err := s.repo.UpdateTask(task); err != nil {
		return Tasks{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func (s *taskService) GetTasksByUserID(userID uint) ([]Tasks, error) {
	return s.repo.GetTasksByUserID(userID)
}
