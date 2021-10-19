package users

import "jwt-in-golang/entity"

type Repository interface {
	GetUser(key string) (entity.User, error)
	AddUser(user entity.User) (err error)
}
