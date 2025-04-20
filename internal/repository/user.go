package repository

import (
	"super-cms/helper"
	"super-cms/internal/entity"

	"github.com/go-stack/stack"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r userRepository) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (r userRepository) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := r.db.Where(entity.User{Email: email}).
		Select("id", "email", "password", "name", "username", "alias", "foto").
		First(&user).Error
	if err != nil {
		r.traceErr(err)
	}
	return user, err
}
