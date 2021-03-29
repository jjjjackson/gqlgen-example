package repository

import (
	"github.com/jjjjackson/gqlgen-example/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll(*gorm.DB) ([]*model.User, error)
	FindOneByEmail(*gorm.DB, string) (*model.User, error)

	Create(*gorm.DB, *model.User) (*model.User, error)
}
