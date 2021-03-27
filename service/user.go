package service

import (
	"github.com/jjjjackson/gqlgen-example/model"
	"github.com/jjjjackson/gqlgen-example/repository"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserService interface {
	ListUsers() ([]*model.User, error)

	CreateUser(*model.User) (*model.User, error)
}

type UserService struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserService(
	db *gorm.DB,
	userRepo repository.UserRepository,
) *UserService {
	return &UserService{
		db:   db,
		repo: userRepo,
	}
}

func (svc *UserService) ListUsers() ([]*model.User, error) {
	return svc.repo.FindAll(svc.db)
}

func (svc *UserService) GetUser(email string) (*model.User, error) {
	return svc.repo.FindOneByEmail(svc.db, email)
}

func (svc *UserService) CreateUser(email, password string) (*model.User, error) {
	user, err := model.CreateUser(email, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err = svc.repo.Create(svc.db, user)

	return user, nil
}
