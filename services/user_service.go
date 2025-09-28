package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/triadynata/project-management/models"
	"github.com/triadynata/project-management/repositories"
	"github.com/triadynata/project-management/utils"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	//kita harus mengecek email yang terdaftar apakah sudah dipakai atau belum
	//hasing password
	//set role
	//simpan user

	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}
	hased, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}
