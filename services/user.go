package services

import (
	"errors"
	"ttk/models"
	"ttk/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUserByID возвращает пользователя по ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser обновляет данные пользователя
func (s *UserService) UpdateUser(id uint, updatedUser *models.User) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Обновляем только разрешенные поля
	if updatedUser.Username != "" {
		user.Username = updatedUser.Username
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}

// DeleteUser удаляет пользователя
func (s *UserService) DeleteUser(id uint) error {
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.userRepo.Delete(id)
}
