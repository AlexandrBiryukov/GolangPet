package Handlers

import (
	"context"
	"golang/internal/userService"
	"golang/internal/web/users"
	"time"
)

type UserHandler struct {
	service userService.Service
}

func NewUserHandler(s userService.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	all, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}

	resp := users.GetUsers200JSONResponse{}

	for _, u := range all {
		var deletedAt *time.Time
		if u.DeletedAt.Valid {
			t := u.DeletedAt.Time
			deletedAt = &t
		}

		resp = append(resp, users.User{
			Id:        uint64(u.ID),
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: deletedAt,
		})
	}

	return resp, nil
}

func (h *UserHandler) PostUsers(_ context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	body := req.Body

	created, err := h.service.Create(body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if created.DeletedAt.Valid {
		t := created.DeletedAt.Time
		deletedAt = &t
	}

	return users.PostUsers201JSONResponse{
		Id:        uint64(created.ID),
		Email:     created.Email,
		Password:  created.Password,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
		DeletedAt: deletedAt,
	}, nil
}

func (h *UserHandler) PatchUserById(_ context.Context, req users.PatchUserByIdRequestObject) (users.PatchUserByIdResponseObject, error) {
	body := req.Body

	updated, err := h.service.PatchByID(uint(req.Id), body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if updated.DeletedAt.Valid {
		t := updated.DeletedAt.Time
		deletedAt = &t
	}

	return users.PatchUserById200JSONResponse{
		Id:        uint64(updated.ID),
		Email:     updated.Email,
		Password:  updated.Password,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
		DeletedAt: deletedAt,
	}, nil
}

func (h *UserHandler) DeleteUserById(_ context.Context, req users.DeleteUserByIdRequestObject) (users.DeleteUserByIdResponseObject, error) {
	if err := h.service.DeleteByID(uint(req.Id)); err != nil {
		return nil, err
	}
	return users.DeleteUserById204Response{}, nil
}
