package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
