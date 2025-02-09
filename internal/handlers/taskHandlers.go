package handlers

import (
	"context"
	"net/http"
	"restapi/internal/tasksService"
	"restapi/internal/web/tasks"

	"github.com/labstack/echo"
)

type Handler struct {
	Service *tasksService.TaskService
}

func NewHandler(service *tasksService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Преобразуем ID из `int` в `uint`
	id := uint(request.Id)

	// Удаляем задачу
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		// Если задача не найдена, возвращаем 404
		return tasks.DeleteTasksId404Response{}, nil
	}

	// Успешное удаление
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Преобразуем ID из int в uint
	id := uint(request.Id)

	// Получаем данные для обновления из тела запроса
	taskRequest := request.Body
	if taskRequest == nil {
		// Если тело запроса отсутствует
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	// Создаем структуру задачи для обновления
	updatedTask := tasksService.Task{
		ID:     id,
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Обновляем задачу через сервис
	updatedTaskResult, err := h.Service.UpdateTaskByID(id, updatedTask)
	if err != nil {
		// Если задача не найдена, возвращаем 404
		return tasks.PatchTasksId404Response{}, nil
	}

	// Формируем успешный ответ
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTaskResult.ID,
		Task:   &updatedTaskResult.Task,
		IsDone: &updatedTaskResult.IsDone,
	}
	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
