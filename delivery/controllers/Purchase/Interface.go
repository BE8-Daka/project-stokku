package controllers

import "github.com/labstack/echo/v4"

type PurchaseController interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
}