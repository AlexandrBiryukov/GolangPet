package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		taskText  string
		isDone    bool
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name:     "успешное создание",
			taskText: "Test",
			isDone:   false,
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.AnythingOfType("*services.Tasks")).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "ошибка репозитория",
			taskText: "Test",
			isDone:   false,
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.AnythingOfType("*services.Tasks")).Return(errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name:      "пустой taskText -> ошибка в сервисе",
			taskText:  "",
			isDone:    false,
			mockSetup: func(m *MockTaskRepository) {}, // repo не должен вызываться
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			svc := NewTaskService(mockRepo)

			res, err := svc.CreateTask(tt.taskText, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.taskText, res.Task)
				assert.Equal(t, tt.isDone, res.IsDone)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTasks(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantLen   int
		wantErr   bool
	}{
		{
			name: "успешно вернул список",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTasks").Return([]Tasks{
					{ID: 1, Task: "A", IsDone: false},
					{ID: 2, Task: "B", IsDone: true},
				}, nil)
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "ошибка репозитория",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTasks").Return(nil, errors.New("db error"))
			},
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			svc := NewTaskService(mockRepo)

			res, err := svc.GetTasks()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, res, tt.wantLen)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestPatchTask(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		taskText  *string
		isDone    *bool
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name:     "успешный патч task",
			id:       "1",
			taskText: ptrStr("New"),
			isDone:   nil,
			mockSetup: func(m *MockTaskRepository) {
				// 1) сначала сервис получит текущую задачу
				m.On("GetTaskById", "1").Return(Tasks{ID: 1, Task: "Old", IsDone: false}, nil)
				// 2) потом обновит
				m.On("UpdateTask", mock.MatchedBy(func(t Tasks) bool {
					return t.ID == 1 && t.Task == "New" && t.IsDone == false
				})).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "успешный патч is_done",
			id:       "1",
			taskText: nil,
			isDone:   ptrBool(true),
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", "1").Return(Tasks{ID: 1, Task: "Old", IsDone: false}, nil)
				m.On("UpdateTask", mock.MatchedBy(func(t Tasks) bool {
					return t.ID == 1 && t.Task == "Old" && t.IsDone == true
				})).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "ошибка на GetTaskById",
			id:       "1",
			taskText: ptrStr("New"),
			isDone:   nil,
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", "1").Return(Tasks{}, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:     "ошибка на UpdateTask",
			id:       "1",
			taskText: ptrStr("New"),
			isDone:   nil,
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", "1").Return(Tasks{ID: 1, Task: "Old", IsDone: false}, nil)
				m.On("UpdateTask", mock.AnythingOfType("services.Tasks")).Return(errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name:     "taskText пустой -> ошибка сервиса",
			id:       "1",
			taskText: ptrStr(""),
			isDone:   nil,
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", "1").Return(Tasks{ID: 1, Task: "Old", IsDone: false}, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			svc := NewTaskService(mockRepo)

			res, err := svc.PatchTask(tt.id, tt.taskText, tt.isDone)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, uint(1), res.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное удаление",
			id:   "1",
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", "1").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении",
			id:   "1",
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", "1").Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			svc := NewTaskService(mockRepo)

			err := svc.DeleteTask(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func ptrStr(v string) *string { return &v }
func ptrBool(v bool) *bool    { return &v }
