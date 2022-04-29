package repository

import (
	"project-stokku/entity"
	"time"

	"gorm.io/gorm"
)

type productModel struct {
	DB *gorm.DB
}

func NewProductModel(db *gorm.DB) *productModel {
	return &productModel{
		DB: db,
	}
}

func (u *productModel) Create(product *entity.Product) error {
	if err := u.DB.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (u *productModel) Get(id string) (*entity.Product, error) {
	var product entity.Product

	if err := u.DB.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}

	u.DB.Model(&product).Association("Purchases").Find(&product.Purchases)
	u.DB.Model(&product).Association("Sales").Find(&product.Sales)

	var qtyPurchase, qtySale uint
	
	for i := 0; i < len(product.Purchases); i++ {
		qtyPurchase += product.Purchases[i].Qty
	}

	for i := 0; i < len(product.Sales); i++ {
		qtySale += product.Sales[i].Qty
	} 
	
	product.Stock = product.Stock + qtyPurchase - qtySale

	return &product, nil
}

func (u *productModel) GetAll() ([]entity.Product, error) {
	var products []entity.Product

	if err := u.DB.Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Find(&products).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(products); i++ {
		u.DB.Model(&products[i]).Association("Purchases").Find(&products[i].Purchases)
		u.DB.Model(&products[i]).Association("Sales").Find(&products[i].Sales)

		var qtyPurchase, qtySale uint

		for j := 0; j < len(products[i].Purchases); j++ {
			qtyPurchase += products[i].Purchases[j].Qty
		}

		for j := 0; j < len(products[i].Sales); j++ {
			qtySale += products[i].Sales[j].Qty
		}

		products[i].Stock = products[i].Stock + qtyPurchase - qtySale
	}

	return products, nil
}