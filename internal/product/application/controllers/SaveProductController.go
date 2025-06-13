package controllers

import (
	"net/http"

	"product/application/dtos"
	"product/application/usecases"
)

type SaveProductController struct {
	SaveProductUseCase *usecases.SaveProductUseCase
}

func (controller *SaveProductController) Handle(w http.ResponseWriter, r *http.Request) error {
	err := controller.SaveProductUseCase.Execute(dtos.Product{})
	if err != nil {
		return err
	}

	w.Write([]byte("{ \"ok\": true }"))
	return nil
}
