package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/jjjjackson/gqlgen-example/util/token"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	saltSize     = 6
	passwordMask = "********"
)

type UserList []User

type User struct {
	UID       string `gorm:"primary_key;column:uid" json:uid`
	Email     string
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func CreateUser(email, password string) (*User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	encryptedPassword, err := encryptPassword(password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &User{
		UID:      uid.String(),
		Email:    email,
		Password: encryptedPassword,
		Salt:     token.GenerateToken(saltSize),
	}, nil
}

func (u *User) MaskPassword() {
	u.Password = passwordMask
}

func encryptPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
