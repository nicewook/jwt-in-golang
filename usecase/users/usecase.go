package users

import "jwt-in-golang/entity"

type Usecase interface {
	SignUp(user entity.User) error
	SignIn(user entity.User) (string, error)
	SayHello(user entity.User) (string, error)
}
