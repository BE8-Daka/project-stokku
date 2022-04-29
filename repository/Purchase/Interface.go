package repository

import "project-stokku/entity"

type PurchaseModel interface {
	Create(purchase *entity.Purchase) error
	Get(id string) (*entity.Purchase, error)
	GetAll() ([]entity.Purchase, error)
}