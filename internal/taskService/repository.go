package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// UpdateTaskByID - обновляет задачу по ID
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	err := r.db.First(&existingTask, id).Error
	if err != nil {
		return Task{}, err
	}

	// Обновляем задачу
	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone
	err = r.db.Save(&existingTask).Error
	return existingTask, err
}

// DeleteTaskByID - удаляет задачу по ID
func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&task).Error
	return err
}
