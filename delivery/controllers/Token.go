package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}

	return 0
}