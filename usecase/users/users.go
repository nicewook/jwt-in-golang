package users

import (
	"fmt"
	"jwt-in-golang/entity"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) SignUp(input entity.User) error {
	// need to hash the password
	hashed, err := HashAndSalt(input.Password)
	if err != nil {
		log.Printf("failed to encrypt password with error: %s", err)
		return err
	}

	log.Printf("username: %v", input.Username)
	details, err := svc.Repository.GetUser(input.Username)
	if err != nil {
		log.Printf("failed to get user with error: %s", err)
		return err
	}

	if details.Username == input.Username {
		return entity.ErrUserAlreadyExists
	}

	// sending the input but only changing the password to the hashed
	input.Password = hashed

	err = svc.Repository.AddUser(input)
	if err != nil {
		log.Printf("failed to add user with error: %s", err)
		return err
	}
	return nil
}

func (svc *Service) SignIn(input entity.User) (token string, err error) {
	details, err := svc.Repository.GetUser(input.Username)
	if err != nil {
		return token, err
	}

	if details.Username == "" {
		return token, entity.ErrUserDoesNotExist
	}

	if !(ComparePasswords(details.Password, input.Password)) {
		return token, entity.ErrInvalidPassword
	}

	token, err = CreateToken(details, svc.JWTSecret)
	if err != nil {
		log.Printf("failed to create token with error: %s", err)
		return token, err
	}
	return token, nil
}

func (ser *Service) SayHello(input entity.User) (message string, err error) {
	details, err := ser.Repository.GetUser(input.Username)
	if err != nil {
		log.Printf("failed to retrieve user data: %s", err)
		return message, err
	}

	message = fmt.Sprintf("Hello %s", details.FirstName)

	return message, nil
}

func HashAndSalt(password string) (string, error) {
	// hash and salt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashed, plain string) bool {
	// compare the hashed and plain passwords
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func CreateToken(user entity.User, jwtSecret string) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"username": user.Username,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour * time.Duration(1)).Unix(),
	}

	log.Printf("token.Claims %+v", token.Claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
