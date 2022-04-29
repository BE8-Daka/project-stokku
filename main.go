package main

import (
	"fmt"
	"log"
	"project-stokku/config"
	pc "project-stokku/delivery/controllers/Product"
	membeli "project-stokku/delivery/controllers/Purchase"
	menjual "project-stokku/delivery/controllers/Sale"
	uc "project-stokku/delivery/controllers/User"
	"project-stokku/delivery/routes"
	pm "project-stokku/repository/Product"
	beli "project-stokku/repository/Purchase"
	jual "project-stokku/repository/Sale"
	um "project-stokku/repository/User"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)

	userModel := um.NewUserModel(db)
	productModel := pm.NewProductModel(db)
	purchaseModel := beli.NewPurchaseModel(db)
	saleModel := jual.NewSaleModel(db)

	userController := uc.NewUserController(userModel, validator.New())
	productController := pc.NewProductController(productModel, validator.New())
	purchaseController := membeli.NewPurchaseController(purchaseModel, validator.New())
	saleController := menjual.NewSaleController(saleModel, validator.New())

	server := echo.New()
	
	routes.RoutePath(server, userController, productController, purchaseController, saleController)
	log.Fatal(server.Start(fmt.Sprintf("%d", conf.Port)))
}
