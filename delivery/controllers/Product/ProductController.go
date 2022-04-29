package controllers

import (
	"net/http"
	controllers "project-stokku/delivery/controllers"
	helpers "project-stokku/delivery/helpers/request"
	views "project-stokku/delivery/helpers/view"
	"project-stokku/entity"
	repository "project-stokku/repository/Product"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productController struct {
	Connect repository.ProductModel
	Validator *validator.Validate
}

func NewProductController(conn repository.ProductModel, valid *validator.Validate) *productController {
	return &productController{
		Connect:  conn,
		Validator: valid,
	}
}

func (pc *productController) Create() echo.HandlerFunc {
	return func (c echo.Context) error {
		var product helpers.ProductRequest
		userID := controllers.ExtractTokenUserId(c)
	
		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, views.BadRequest("Invalid argument exception"))
		}
	
		if err := pc.Validator.Struct(&product); err != nil {
			return c.JSON(http.StatusBadRequest, views.InputRequest("is Required", "name"))
		}
	
		pc.Connect.Create(&entity.Product{
			Name: product.Name,
			UserID: uint(int(userID)),
		})
	
		return c.JSON(http.StatusCreated, views.SuccessInsert("Create", "Product", product))
	}
}

func (u *productController) Get(c echo.Context) error {
	id := c.Param("id")
	product, err := u.Connect.Get(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Product"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Product", product))
}

func (u *productController) GetAll(c echo.Context) error {
	products, err := u.Connect.GetAll()
	
	if err != nil {
		return err
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Products"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Products", products))
}