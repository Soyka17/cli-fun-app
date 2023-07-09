package db

import "ExperienceBank/internal/entity"

type DatabaseRepository interface {
	SaveNewUser(*entity.User) error
	GetUserById(int) (*entity.User, error)
	GetNewId() int
	UpdateUserInfo(user *entity.User) error
}
