package tasksService

import "errors"

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) PostTask(task Task) (Task, error) {
	if task.UserID == 0 {
		return Task{}, errors.New("user_id is required")
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUser(userID)
}
