package repository

import (
	"project-stokku/entity"
	"time"

	"gorm.io/gorm"
)

type purchaseModel struct {
	DB *gorm.DB
}

func NewPurchaseModel(db *gorm.DB) *purchaseModel {
	return &purchaseModel{
		DB: db,
	}
}

func (u *purchaseModel) Create(purchase *entity.Purchase) error {
	var product entity.Product

	if err := u.DB.Where("id = ?", purchase.ProductID).Find(&product).Error; err != nil {
		return err
	}

	if product.ID != purchase.ProductID {
		return gorm.ErrRecordNotFound
	}

	if err := u.DB.Create(purchase).Error; err != nil {
		return err
	}

	u.DB.Create(entity.PurchaseProduct{
		PurchaseID: purchase.ID,
		ProductID: purchase.ProductID,
	})

	return nil
}

func (u *purchaseModel) Get(id string) (*entity.Purchase, error) {
	var purchase entity.Purchase

	if err := u.DB.Where("id = ?", id).Find(&purchase).Error; err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (u *purchaseModel) GetAll() ([]entity.Purchase, error) {
	var purchases []entity.Purchase

	if err := u.DB.Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}