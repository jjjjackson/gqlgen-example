package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jjjjackson/gqlgen-example/util"

	"github.com/pkg/errors"
)

const (
	saltSize     = 6
	passwordMask = "********"
)

type UserList []User

type User struct {
	UID       string `gorm:"primary_key;column:uid" json:"uid"`
	Email     string
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func NewUser(email, password string) (*User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	encryptedPassword, err := util.EncryptPassword(password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &User{
		UID:      uid.String(),
		Email:    email,
		Password: encryptedPassword,
		Salt:     util.RandomString(saltSize),
	}, nil
}

func (u *User) MaskPassword() {
	u.Password = passwordMask
}
