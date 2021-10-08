package handler

import (
	"fmt"
	"jwt-in-golang/api/middleware"
	"jwt-in-golang/entity"
	"jwt-in-golang/model"
	"jwt-in-golang/usecase/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Usecase users.UseCase
}

func NewGinHandler(usecase users.UseCase, jwtSecret string) *gin.Engine {
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

func (h *GinHandler) signIn(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		errs := []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeValidation,
				ErrMsg:  err.Error(),
			},
		}
		badRequest(c, http.StatusBadRequest, "Invalid request", errs)
	}

	token, err := h.Usecase.SignIn(user)
	if err != nil {
		ErrHandler(err, c)
	}

	res := SigninOutput{Token: token}
	c.JSON(http.StatusOK, res)
}

func (h *GinHandler) signUp(c *gin.Context) {
	var signupObj entity.User

	if err := c.ShouldBindJSON(&signupObj); err != nil {
		errs := []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeValidation,
				ErrMsg:  err.Error(),
			},
		}
		badRequest(c, http.StatusBadRequest, "Invalid request", errs)
	}

	if err := h.Usecase.SignUp(signupObj); err != nil { // real databbase save
		ErrHandler(err, c)
	}

	res := SignupOutput{Message: fmt.Sprintf("user %s created successfully", signupObj.Username)}
	c.JSON(http.StatusCreated, res)
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

func ok(c *gin.Context, status int, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, model.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}

func badRequest(c *gin.Context, status int, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}
