package userService

import (
	"errors"
	"strings"
)

type Service interface {
	Create(email, password string) (User, error)
	GetAll() ([]User, error)
	PatchByID(id uint, email, password *string) (User, error)
	DeleteByID(id uint) error
	GetTasksForUser(userID uint) ([]Task, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(email, password string) (User, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return User{}, errors.New("Требуется адрес электронной почты")
	}
	if password == "" {
		return User{}, errors.New("требуется пароль")
	}

	u := User{Email: email, Password: password}
	if err := s.repo.Create(&u); err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *service) PatchByID(id uint, email, password *string) (User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return User{}, err
	}

	if email != nil {
		e := strings.TrimSpace(*email)
		if e == "" {
			return User{}, errors.New("Адрес электронной почты не может быть пустым")
		}
		u.Email = e
	}
	if password != nil {
		if *password == "" {
			return User{}, errors.New("Пароль не может быть пустым")
		}
		u.Password = *password
	}

	if err := s.repo.Update(&u); err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) DeleteByID(id uint) error {
	return s.repo.DeleteByID(id)
}

func (s *service) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUser(userID)
}
