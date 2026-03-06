package Handlers

import (
	"context"
	"golang/internal/services"
	"golang/internal/web/tasks"
	"strconv"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(s services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetTasks()
	if err != nil {
		return nil, err
	}

	resp := tasks.GetTasks200JSONResponse{}
	for _, t := range allTasks {
		resp = append(resp, tasks.Task{
			Id:     uint64(t.ID),
			Task:   t.Task,
			IsDone: t.IsDone,
			UserId: uint64(t.UserID),
		})
	}
	return resp, nil
}

func (h *TaskHandler) GetTasksByUserID(_ context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	allTasks, err := h.service.GetTasksByUserID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	resp := tasks.GetTasksByUserID200JSONResponse{}
	for _, t := range allTasks {
		resp = append(resp, tasks.Task{
			Id:     uint64(t.ID),
			Task:   t.Task,
			IsDone: t.IsDone,
			UserId: uint64(t.UserID),
		})
	}
	return resp, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body

	created, err := h.service.CreateTask(body.Task, body.IsDone, uint(body.UserId))
	if err != nil {
		return nil, err
	}

	return tasks.PostTasks201JSONResponse{
		Id:     uint64(created.ID),
		Task:   created.Task,
		IsDone: created.IsDone,
	}, nil
}

func (h *TaskHandler) PatchTaskById(_ context.Context, request tasks.PatchTaskByIdRequestObject) (tasks.PatchTaskByIdResponseObject, error) {
	idStr := strconv.FormatUint(request.Id, 10)
	body := request.Body // *TaskPatch (поля указатели)

	updated, err := h.service.PatchTask(idStr, body.Task, body.IsDone)
	if err != nil {
		return nil, err
	}

	return tasks.PatchTaskById200JSONResponse{
		Id:     uint64(updated.ID),
		Task:   updated.Task,
		IsDone: updated.IsDone,
		UserId: uint64(updated.UserID),
	}, nil
}

func (h *TaskHandler) DeleteTaskById(_ context.Context, request tasks.DeleteTaskByIdRequestObject) (tasks.DeleteTaskByIdResponseObject, error) {
	idStr := strconv.FormatUint(request.Id, 10)

	if err := h.service.DeleteTask(idStr); err != nil {
		return nil, err
	}

	return tasks.DeleteTaskById204Response{}, nil
}
