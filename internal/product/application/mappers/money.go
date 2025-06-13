package mappers

import (
	"product/domain/valueobjects"
)

func FromStringToMoney(priceString string) *valueobjects.Money {
	return &valueobjects.Money{
		Value:    1,
		Currency: valueobjects.Currency{},
	}
}
