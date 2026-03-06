package services

import "gorm.io/gorm"

type TaskRepositoty interface {
	CreateTask(task *Tasks) error
	GetTasks() ([]Tasks, error)
	GetTaskById(taskId string) (Tasks, error)
	UpdateTask(task Tasks) error
	DeleteTask(id string) error
	GetTasksByUserID(userID uint) ([]Tasks, error)
}
type taskRepositoty struct {
	db *gorm.DB
}

func NewTaskRepositoty(db *gorm.DB) TaskRepositoty {
	return &taskRepositoty{db: db}
}

func (r *taskRepositoty) CreateTask(task *Tasks) error {
	return r.db.Create(task).Error
}

func (r *taskRepositoty) GetTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepositoty) GetTaskById(taskId string) (Tasks, error) {
	var task Tasks
	err := r.db.First(&task, "id = ?", taskId).Error
	return task, err

}
func (r *taskRepositoty) UpdateTask(task Tasks) error {
	return r.db.Save(&task).Error
}
func (r *taskRepositoty) DeleteTask(id string) error {
	return r.db.Delete(&Tasks{}, "id = ?", id).Error
}

func (r *taskRepositoty) GetTasksByUserID(userID uint) ([]Tasks, error) {
	var tasks []Tasks
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
