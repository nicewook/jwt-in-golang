package main

import (
	"jwt-in-golang/handler"
	"jwt-in-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	api := r.Group("/api")
	{
		api.Use(middleware.ValidateToken())

		product := api.Group("/product")
		{
			product.Use(middleware.Authorization([]int{1}))

			product.GET("/", handler.GetAll)
			product.POST("/", middleware.Authorization([]int{4}), handler.AddProduct)
		}

		user := api.Group("/user")
		{
			user.GET("/", handler.GetUser)
		}
	}
	return r
}

func main() {
	server := NewServer()
	server.Run()
}
