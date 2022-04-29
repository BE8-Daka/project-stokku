package repository

import "project-stokku/entity"

type ProductModel interface {
	Create(product *entity.Product) error
	Get(id string) (*entity.Product, error)
	GetAll() ([]entity.Product, error)
}