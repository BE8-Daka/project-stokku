package helpers

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required"`
}

type ProductRequest struct {
	Name string `json:"name" validate:"required"`
}

type PurchaseRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Price     uint `json:"price" validate:"required"`
	Qty       uint `json:"qty" validate:"required"`
}

type SaleRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Price     uint `json:"price" validate:"required"`
	Qty       uint `json:"qty" validate:"required"`
}