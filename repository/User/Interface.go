package repository

import "project-stokku/entity"

type UserModel interface {
	CheckDuplicate(email string) bool
	Create(user *entity.User) error
	Login(email, password string) (bool, entity.User)
}