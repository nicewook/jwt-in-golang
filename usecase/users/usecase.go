package users

import "jwt-in-golang/entity"

type UseCase interface {
	SignUp(user entity.User) error
	SignIn(user entity.User) (string, error)
	SayHello(user entity.User) (string, error)
}
