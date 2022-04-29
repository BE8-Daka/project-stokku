package controllers

import "github.com/labstack/echo/v4"

type ProductController interface {
	Create() echo.HandlerFunc
	Get(c echo.Context) error
	GetAll(c echo.Context) error
}