package userService

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetByID(id uint) (User, error)
	Update(user *User) error
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetByID(id uint) (User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return u, err
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) DeleteByID(id uint) error {
	res := r.db.Delete(&User{}, id)
	if res.Error != nil {
		return res.Error
	}
	// если не нашли — считаем ошибкой (можно иначе)
	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
