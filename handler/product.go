package handler

import (
	"jwt-in-golang/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetAll(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, []Product{
		{
			ID:   1,
			Name: "Product 1",
		},
		{
			ID:   2,
			Name: "Product 2",
		},
	})
}
func AddProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errs := []model.ErrorDetail{
			{
				ErrType: model.ErrorTypeValidation,
				ErrMsg:  err.Error(),
			},
		}
		badRequest(c, http.StatusBadRequest, "Invalid request", errs)
	}
	product.ID = 10
	ok(c, http.StatusCreated, "Product added", product)
}
