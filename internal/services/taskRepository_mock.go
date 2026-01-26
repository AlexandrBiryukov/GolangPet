package services

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository - поддельный репозиторий
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *Tasks) error {
	args := m.Called(task)

	return args.Error(0)
}

func (m *MockTaskRepository) GetTasks() ([]Tasks, error) {
	args := m.Called()
	var tasks []Tasks
	if res := args.Get(0); res != nil {
		tasks = res.([]Tasks)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task Tasks) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTaskById(taskId string) (Tasks, error) {
	args := m.Called(taskId)
	var task Tasks
	if res := args.Get(0); res != nil {
		task = res.(Tasks)
	}
	return task, args.Error(1)
}
