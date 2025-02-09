package handlers

import (
	"net/http"
	"restapi/internal/web/users"
	"restapi/usersService"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	Service *usersService.UserService
}

// GetUsers реализует users.ServerInterface.
func (h *UserHandlers) GetUsers(ctx echo.Context) error {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка получения пользователей")
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		tempID := int(usr.ID) // Преобразуем uint в int
		id := &tempID

		user := users.User{
			Id:        id,
			Email:     &usr.Email,
			Password:  &usr.Password,
			DeleteAt:  usr.DeletedAt,
			CreatedAt: &usr.CreatedAt,
			UpdateAt:  &usr.UpdatedAt,
		}

		response = append(response, user)
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostUsers реализует users.ServerInterface (создание пользователя).
func (h *UserHandlers) PostUsers(ctx echo.Context) error {
	var newUser users.User
	if err := ctx.Bind(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Неверный формат запроса")
	}

	createdUser, err := h.Service.PostUser(usersService.User{
		Email:    *newUser.Email,
		Password: *newUser.Password,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка создания пользователя")
	}

	// Преобразование *uint в *int
	tempID := int(createdUser.ID)
	id := &tempID

	response := users.PostUsers201JSONResponse{
		Id:        id,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		DeleteAt:  createdUser.DeletedAt,
		CreatedAt: &createdUser.CreatedAt,
		UpdateAt:  &createdUser.UpdatedAt,
	}

	return ctx.JSON(http.StatusCreated, response)
}

// PatchUsersId реализует users.ServerInterface (обновление пользователя).
func (h *UserHandlers) PatchUsersId(ctx echo.Context, id int) error {
	// Получаем тело запроса
	var updateRequest users.PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Неверный формат запроса")
	}

	// Преобразуем `id` из `int` в `uint`
	userID := uint(id)

	// Формируем структуру для обновления
	updatedUser := usersService.User{
		ID: userID,
	}

	// Заполняем `Email` только если он передан в запросе
	if updateRequest.Email != nil {
		updatedUser.Email = *updateRequest.Email
	}

	// Заполняем `Password` только если он передан в запросе
	if updateRequest.Password != nil {
		updatedUser.Password = *updateRequest.Password
	}

	// Вызываем сервис обновления
	updatedUserResult, err := h.Service.PatchUserByID(userID, updatedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Пользователь не найден")
	}

	// Преобразуем `uint` в `int` для корректного JSON-ответа
	tempID := int(updatedUserResult.ID)
	idPtr := &tempID

	response := users.PatchUsersId200JSONResponse{
		Id:        idPtr,
		Email:     &updatedUserResult.Email,
		Password:  &updatedUserResult.Password,
		DeleteAt:  updatedUserResult.DeletedAt,
		CreatedAt: &updatedUserResult.CreatedAt,
		UpdateAt:  &updatedUserResult.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteUsersId реализует users.ServerInterface (удаление пользователя).
func (h *UserHandlers) DeleteUsersId(ctx echo.Context, id int) error {
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Пользователь не найден")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// Используем Handler при создании экземпляра
func NewUserHandler(service *usersService.UserService) *UserHandlers {
	return &UserHandlers{Service: service}
}
