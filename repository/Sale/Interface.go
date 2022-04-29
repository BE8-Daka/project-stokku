package repository

import "project-stokku/entity"

type SaleModel interface {
	Create(sale *entity.Sale) error
	Get(id string) (*entity.Sale, error)
	GetAll() ([]entity.Sale, error)
}