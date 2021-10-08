package handler

import (
	"jwt-in-golang/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
