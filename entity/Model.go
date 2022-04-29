package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Products []Product `gorm:"foreignkey:UserID"`
	Purchases []Purchase `gorm:"foreignkey:UserID"`
	Sales []Sale `gorm:"foreignkey:UserID"`
}

type Product struct {
	gorm.Model
	Name string
	Stock uint
	UserID uint
	Purchases []Purchase `gorm:"many2many:purchase_products"`
	Sales []Sale `gorm:"many2many:sale_products"`
}

type Purchase struct {
	gorm.Model
	ProductID uint
	UserID uint
	Price uint
	Qty uint
}

type Sale struct {
	gorm.Model
	ProductID uint
	Price uint
	UserID uint
	Qty uint
}

type PurchaseProduct struct {
	ProductID uint
	PurchaseID uint
}

type SaleProduct struct {
	ProductID uint
	SaleID uint
}