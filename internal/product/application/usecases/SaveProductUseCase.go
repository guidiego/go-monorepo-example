package usecases

import (
	"product/application/dtos"
	"product/application/mappers"
	"product/domain/entities"
	"product/domain/repositories"
)

type SaveProductUseCase struct {
	Repository repositories.IProductRepository
}

func (usecase *SaveProductUseCase) Execute(productDto dtos.Product) error {
	product := entities.Product{
		Id:          "",
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       mappers.FromStringToMoney(productDto.Price),
	}

	err := product.Validate()
	if err != nil {
		return err
	}

	usecase.Repository.Save(&product)
	return nil
}
