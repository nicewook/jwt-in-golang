package entity

import "errors"

var ErrUserAlreadyExists = errors.New("username already exists")

var ErrUserDoesNotExist = errors.New("username does not exist")

var ErrInvalidPassword = errors.New("invalid password")

var ErrInvalidInput = errors.New("invalid input")
