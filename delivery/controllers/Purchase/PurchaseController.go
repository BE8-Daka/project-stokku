package controllers

import (
	"net/http"
	controllers "project-stokku/delivery/controllers"
	helpers "project-stokku/delivery/helpers/request"
	views "project-stokku/delivery/helpers/view"
	"project-stokku/entity"
	repository "project-stokku/repository/Purchase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type purchaseController struct {
	Connect repository.PurchaseModel
	Validator *validator.Validate
}

func NewPurchaseController(conn repository.PurchaseModel, valid *validator.Validate) *purchaseController {
	return &purchaseController{
		Connect:  conn,
		Validator: valid,
	}
}

func (pc *purchaseController) Create(c echo.Context) error {
	var purchase helpers.PurchaseRequest
	userID := controllers.ExtractTokenUserId(c)

	if err := c.Bind(&purchase); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, views.BadRequest("Invalid argument exception"))
	}

	if err := pc.Validator.Struct(&purchase); err != nil {
		return c.JSON(http.StatusBadRequest, views.InputRequest("is Required","product_id, price, qty"))
	}

	PurchaseNew := entity.Purchase{
		ProductID: purchase.ProductID,
		Price: purchase.Price,
		Qty: purchase.Qty,
		UserID: uint(int(userID)),
	}
		
	err := pc.Connect.Create(&PurchaseNew)

	if err != nil {
		return c.JSON(http.StatusBadRequest, views.BadRequest(err.Error()))
	}

	return c.JSON(http.StatusCreated, views.SuccessInsert("Menambah", "Stok", purchase))
}

func (u *purchaseController) Get(c echo.Context) error {
	id := c.Param("id")
	purchase, err := u.Connect.Get(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Purchase"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Product yang Dibeli", purchase))
}

func (u *purchaseController) GetAll(c echo.Context) error {
	purchases, err := u.Connect.GetAll()
	
	if err != nil {
		return err
	}

	if len(purchases) == 0 {
		return c.JSON(http.StatusNotFound, views.NotFoundResponse("Purchases"))
	}

	return c.JSON(http.StatusOK, views.SuccessResponse("Get", "Purchases", purchases))
}