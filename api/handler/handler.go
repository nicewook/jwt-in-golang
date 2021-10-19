package handler

import (
	"fmt"
	"jwt-in-golang/api/middleware"
	"jwt-in-golang/entity"
	"jwt-in-golang/usecase/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Usecase users.Usecase
	// Can add ProductUsecase products.Usecase
}

func NewGinEngine(usecase users.Usecase, jwtSecret string) *gin.Engine {
	h := &GinHandler{
		Usecase: usecase,
	}

	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/signin", h.signIn)
		users.POST("/signup", h.signUp)
	}

	r.GET("/hello", middleware.TokenAuthMiddleware(jwtSecret), h.sayHello)
	return r
}

func (h *GinHandler) signUp(c *gin.Context) {
	var signupObj entity.User

	if err := c.ShouldBindJSON(&signupObj); err != nil {
		ErrHandler(err, c)
	}

	if err := h.Usecase.SignUp(signupObj); err != nil { // real databbase save
		ErrHandler(err, c)
	}

	res := SignupOutput{
		Message: fmt.Sprintf("user %s created successfully", signupObj.Username),
	}
	c.JSON(http.StatusCreated, res)
}

func (h *GinHandler) signIn(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		ErrHandler(err, c)
	}

	token, err := h.Usecase.SignIn(user)
	if err != nil {
		ErrHandler(err, c)
	}

	res := SigninOutput{Token: token}
	c.JSON(http.StatusOK, res)
}

func (h *GinHandler) sayHello(c *gin.Context) {
	var l entity.User
	l.Username = c.GetString("username")

	message, err := h.Usecase.SayHello(l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	res := HelloOutput{Message: message}

	c.JSON(200, res)
}
