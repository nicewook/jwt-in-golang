package handler

import (
	"net/http"
	"time"

	"jwt-in-golang/internal/model"
	"jwt-in-golang/internal/token"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginObj model.LoginRequest

	if err := c.ShouldBindJSON(&loginObj); err != nil {
		errs := []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeValidation,
				ErrMsg:  err.Error(),
			},
		}
		badRequest(c, http.StatusBadRequest, "Invalid request", errs)
	}

	// validation
	var claims = &model.JWTClaims{}
	claims.CompanyID = "CompanyID"
	claims.Username = loginObj.Username
	claims.Roles = []int{1, 2, 3}
	claims.Audience = c.Request.Header.Get("referer")

	var tokenIssuedTime = time.Now().UTC()
	var tokenExpirationTime = tokenIssuedTime.Add(2 * time.Hour)

	tokenString, err := token.GenerateToken(claims, tokenExpirationTime)
	if err != nil {
		badRequest(c, http.StatusBadRequest, "Error in generating token", []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeError,
				ErrMsg:  err.Error(),
			},
		})
	}
	ok(c, http.StatusOK, "Token created", tokenString)
}
