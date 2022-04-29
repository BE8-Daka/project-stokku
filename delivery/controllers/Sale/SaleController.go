package controllers

import (
	"net/http"
	controllers "project-stokku/delivery/controllers"
	helpers "project-stokku/delivery/helpers/request"
	views "project-stokku/delivery/helpers/view"
	"project-stokku/entity"
	repository "project-stokku/repository/Sale"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type saleController struct {
	Connect repository.SaleModel
	Validator *validator.Validate
}

func NewSaleController(conn repository.SaleModel, valid *validator.Validate) *saleController {
	return &saleController{
		Connect:  conn,
		Validator: valid,
	}
}

func (pc *saleController) Create(c echo.Context) error {
	var sale helpers.SaleRequest
	userID := controllers.ExtractTokenUserId(c)

	if err := c.Bind(&sale); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, views.BadRequest("Invalid argument exception"))
	}

	if err := pc.Validator.Struct(&sale); err != nil {
		return c.JSON(http.StatusBadRequest, views.InputRequest("is Required","product_id, price, qty"))
	}

	SaleNew := entity.Sale{
		ProductID: sale.ProductID,
		Price: sale.Price,
		Qty: sale.Qty,
		UserID: uint(int(userID)),
	}
	
	err := pc.Connect.Create(&SaleNew)

	if err != nil {
		return c.JSON(http.StatusBadRequest, views.BadRequest(err.Error()))
	}

	return c.JSON(http.StatusCreated, views.SuccessInsert("Menjual", "Product", sale))
}

func (u *saleController) Get(c echo.Context) error {
	id := c.Param("id")
	sale, err := u.Connect.Get(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Purchase"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Product yang Dijual", sale))
}

func (u *saleController) GetAll(c echo.Context) error {
	sales, err := u.Connect.GetAll()
	
	if err != nil {
		return err
	}

	if len(sales) == 0 {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Sales"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Sales", sales))
}