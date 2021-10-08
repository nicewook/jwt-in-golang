package middleware

import (
	"log"
	"net/http"

	"jwt-in-golang/internal/model"
	"jwt-in-golang/internal/token"

	"github.com/gin-gonic/gin"
)

// ReturnUnauthrized is just http response of abort
func ReturnUnauthrized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Error: []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeUnauthorized,
				ErrMsg:  "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}

func ValidateToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("apikey")
		referer := c.Request.Header.Get("referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthrized(c)
		}
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})

		}
		c.Keys["CommandID"] = claims.CompanyID
		c.Keys["Username"] = claims.Username
		c.Keys["Roles"] = claims.Roles
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			ReturnUnauthrized(c)
		}

		rolesVal, ok := c.Keys["Roles"]
		if !ok || rolesVal == nil {
			ReturnUnauthrized(c)
		}
		log.Println("rolesVal:", rolesVal)

		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, role := range roles {
			validation[role] = 0
		}

		for _, validRole := range validRoles {
			if _, ok := validation[validRole]; !ok {
				ReturnUnauthrized(c)
			}
		}
	}
}
