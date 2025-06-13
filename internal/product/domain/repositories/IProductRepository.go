package repositories

import "product/domain/entities"

type IProductRepository interface {
	List() []*entities.Product
	Save(product *entities.Product)
}
