package model

type LoginRequest struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" form:"rememberMe"`
}
