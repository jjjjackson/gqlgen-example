package usecase

import (
	"github.com/jjjjackson/gqlgen-example/domain/model"
	"github.com/jjjjackson/gqlgen-example/domain/repository"
	"github.com/jjjjackson/gqlgen-example/util"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserService struct {
	db       *gorm.DB
	userRepo repository.IUserRepository
}

func NewUserService(
	db *gorm.DB,
	userRepo repository.IUserRepository,
) *UserService {
	return &UserService{
		db:       db,
		userRepo: userRepo,
	}
}

func (svc *UserService) ListUsers() ([]*model.User, error) {
	return svc.userRepo.FindAll(svc.db)
}

func (svc *UserService) GetUser(email string) (*model.User, error) {
	return svc.userRepo.FindOneByEmail(svc.db, email)
}

func (svc *UserService) CreateUser(email, password string) (*model.User, error) {
	user, err := model.NewUser(email, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err = svc.userRepo.Create(svc.db, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func (svc *UserService) Login(email, password string) (*model.User, error) {

	user, err := svc.userRepo.FindOneByEmail(svc.db, email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if util.CheckPasswordHash(password, user.Password) {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
