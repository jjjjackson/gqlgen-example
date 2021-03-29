package datastore

import (
	"gorm.io/gorm"

	"github.com/jjjjackson/gqlgen-example/domain/model"
	"github.com/pkg/errors"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repo UserRepository) FindAll(db *gorm.DB) ([]*model.User, error) {
	users := []*model.User{}

	db.Find(&users)

	for _, user := range users {
		user.MaskPassword()
	}

	return users, nil
}

func (repo UserRepository) FindOneByEmail(db *gorm.DB, email string) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func (repo UserRepository) Create(db *gorm.DB, user *model.User) (*model.User, error) {
	db.Create(&user)
	user.MaskPassword()

	return user, nil
}
