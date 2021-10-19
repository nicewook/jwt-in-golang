package main

import (
	"jwt-in-golang/api/handler"
	"jwt-in-golang/repository/sqlite"
	"jwt-in-golang/usecase/users"
	"os"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-jwt-secret"
	}

	repo := sqlite.NewSQLiteDB()
	svc := users.LoadService(repo, jwtSecret)

	r := handler.NewGinEngine(svc, jwtSecret)
	return r
}

func main() {
	server := NewServer()
	server.Run()
}
