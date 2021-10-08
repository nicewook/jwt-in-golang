package main

import (
	"jwt-in-golang/api/handler"
	"jwt-in-golang/repository/sqlite"
	"jwt-in-golang/usecase/users"
	"log"
	"os"

	"github.com/dwdcth/consoleEx"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
)

func NewServer() *gin.Engine {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.CallerSkipFrameCount = 2
	// logger := log.With().Logger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	out := consoleEx.ConsoleWriterEx{Out: colorable.NewColorableStdout()}
	logger := zerolog.New(out).With().Caller().Timestamp().Logger()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-jwt-secret"
	}

	repo := sqlite.NewSQLiteDB()
	srv := users.LoadService(repo, &logger)

	// gin.SetMode(gin.ReleaseMode)
	r := handler.NewGinHandler(srv, secret)

	// r := gin.Default()
	// r.POST("/login", handler.LoginHandler)

	// api := r.Group("/api")
	// {
	// 	api.Use(middleware.ValidateToken())

	// 	product := api.Group("/product")
	// 	{
	// 		product.Use(middleware.Authorization([]int{1}))

	// 		product.GET("/", handler.GetAll)
	// 		product.POST("/", middleware.Authorization([]int{4}), handler.AddProduct)
	// 	}

	// 	user := api.Group("/user")
	// 	{
	// 		user.GET("/", handler.GetUser)
	// 	}
	// }
	return r
}

func main() {
	server := NewServer()
	server.Run()
}
