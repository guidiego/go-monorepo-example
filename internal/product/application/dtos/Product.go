package dtos

type Product struct {
	Id          string `json:"id"`
	Price       string `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
