package services

import "gorm.io/gorm"

type TaskRepositoty interface {
	CreateTask(task RequestBody) error
	GetTasks() ([]RequestBody, error)
	GetTaskById(taskId string) (RequestBody, error)
	UpdateTask(task RequestBody) error
	DeleteTask(id string) error
}
type taskRepositoty struct {
	db *gorm.DB
}

func NewTaskRepositoty(db *gorm.DB) TaskRepositoty {
	return &taskRepositoty{db: db}
}

func (r *taskRepositoty) CreateTask(task RequestBody) error {
	return r.db.Create(&task).Error
}

func (r *taskRepositoty) GetTasks() ([]RequestBody, error) {
	var tasks []RequestBody
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepositoty) GetTaskById(taskId string) (RequestBody, error) {
	var task RequestBody
	err := r.db.First(&task, "id = ?", taskId).Error
	return task, err

}
func (r *taskRepositoty) UpdateTask(task RequestBody) error {
	return r.db.Save(&task).Error
}
func (r *taskRepositoty) DeleteTask(id string) error {
	return r.db.Delete(&RequestBody{}, "id = ?", id).Error
}
