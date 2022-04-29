package repository

import (
	"project-stokku/entity"
	"time"

	"gorm.io/gorm"
)

type saleModel struct {
	DB *gorm.DB
}

func NewSaleModel(db *gorm.DB) *saleModel {
	return &saleModel{
		DB: db,
	}
}

func (u *saleModel) Create(sale *entity.Sale) error {
	var product entity.Product

	if err := u.DB.Where("id = ? ", sale.ProductID).Find(&product).Error; err != nil {
		return err
	}
	
	if product.ID != sale.ProductID {
		return gorm.ErrRecordNotFound
	}

	if err := u.DB.Create(sale).Error; err != nil {
		return err
	}
	
	u.DB.Create(entity.SaleProduct{
		SaleID: sale.ID,
		ProductID: sale.ProductID,
	})

	return nil
}

func (u *saleModel) Get(id string) (*entity.Sale, error) {
	var sale entity.Sale

	if err := u.DB.Where("id = ?", id).Find(&sale).Error; err != nil {
		return nil, err
	}

	return &sale, nil
}

func (u *saleModel) GetAll() ([]entity.Sale, error) {
	var sales []entity.Sale

	if err := u.DB.Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Find(&sales).Error; err != nil {
		return nil, err
	}

	return sales, nil
}