package mappers

import (
	"product/application/dtos"
	"product/domain/entities"
)

func FromProductDomainToDTO(product *entities.Product) *dtos.Product {
	return &dtos.Product{
		Id:          product.Id,
		Price:       product.Price.FormatString(),
		Name:        product.Name,
		Description: product.Description,
	}
}
