package usersService

import "restapi/internal/tasksService"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) PostUser(user User) (User, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PatchUserByID(id uint, user User) (User, error) {
	return s.repo.PatchUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]tasksService.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
