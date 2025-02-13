package usersService

import (
	"restapi/internal/tasksService"

	"gorm.io/gorm"
)

type UserRepository interface {
	// Передаем в функцию user типа User из orm.go
	// cоздать нового пользователя
	PostUser(user User) (User, error)
	// Ввыводит всех пользователей
	GetUsers() ([]User, error)
	// Отредактировать поля user по его ID
	PatchUserByID(id uint, User User) (User, error)
	// Delete user by ID
	DeleteUserByID(id uint) error
	// GetTasksByUserID —  для получения всех задач конкретного пользователя.
	GetTasksByUserID(userID uint) ([]tasksService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// (r *userRepository) привязывает данную функцию к нашему репозиторию
func (r *userRepository) PostUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetUsers() ([]User, error) {
	var Users []User
	err := r.db.Find(&Users).Error
	return Users, err
}

// PatchUserByID - обновляет пользователя по ID
func (r *userRepository) PatchUserByID(id uint, user User) (User, error) {
	var existingUser User

	// Найти пользователя в базе данных
	err := r.db.First(&existingUser, id).Error
	if err != nil {
		return User{}, err
	}

	// Обновляем только переданные поля
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}

	// Сохраняем обновленного пользователя
	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil
}

// DeleteUserByID - удаляет задачу по ID
func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	return err
}

// GetTasksByUserID —  для получения всех задач конкретного пользователя.
func (r *userRepository) GetTasksByUserID(userID uint) ([]tasksService.Task, error) {
	var tasks []tasksService.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
