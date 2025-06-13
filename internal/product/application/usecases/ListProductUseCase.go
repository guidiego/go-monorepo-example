package usecases

import (
	"product/application/dtos"
	"product/application/mappers"
	"product/domain/repositories"
)

type ListProductUseCase struct {
	Repository repositories.IProductRepository
}

func (usecase *ListProductUseCase) Execute() []*dtos.Product {
	products := usecase.Repository.List()
	productsDtos := make([]*dtos.Product, 0, len(products))
	for _, product := range products {
		productsDtos = append(productsDtos, mappers.FromProductDomainToDTO(product))
	}

	return productsDtos
}
