package controllers

import "github.com/labstack/echo/v4"

type SaleController interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
}