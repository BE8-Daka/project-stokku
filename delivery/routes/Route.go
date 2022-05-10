package routes

import (
	pc "project-stokku/delivery/controllers/Product"
	purchase "project-stokku/delivery/controllers/Purchase"
	sale "project-stokku/delivery/controllers/Sale"
	uc "project-stokku/delivery/controllers/User"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePath(e *echo.Echo, user uc.UserController, product pc.ProductController, beli purchase.PurchaseController, jual sale.SaleController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	
	e.Use(middleware.CORS())
	
	e.POST("/register", user.Create)
	e.POST("/login", user.Login)
	
	auth := e.Group("/admin", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	
	auth.POST("/products", product.Create())
	auth.POST("/purchases", beli.Create)
	auth.POST("/sales", jual.Create)
	
	auth.GET("/history/products/:id", product.Get)
	auth.GET("/history/purchases/:id", beli.Get)
	auth.GET("/history/sales/:id", jual.Get)

	auth.GET("/history/products", product.GetAll)
	auth.GET("/history/purchases", beli.GetAll)
	auth.GET("/history/sales", jual.GetAll)
}
