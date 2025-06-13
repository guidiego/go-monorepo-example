package app

import (
	"github.com/labstack/echo/v4"
	repositories "product/adapters/repositories/inmemory"
	"product/application/controllers"
	"product/application/usecases"
)

type (
	Controllers struct {
		ListProduct *controllers.ListProductController
		SaveProduct *controllers.SaveProductController
	}

	Dependencies struct {
		Echo *echo.Echo
		*Controllers
	}
)

func InitDependencies() *Dependencies {
	inMemoProductRepo := repositories.InMemoryProductRepository{}

	listProductUseCase := usecases.ListProductUseCase{Repository: &inMemoProductRepo}
	saveProductUseCase := usecases.SaveProductUseCase{Repository: &inMemoProductRepo}

	listProductController := controllers.ListProductController{
		ListProductUseCase: &listProductUseCase,
	}
	saveProductController := controllers.SaveProductController{
		SaveProductUseCase: &saveProductUseCase,
	}

	return &Dependencies{
		Echo: echo.New(),
		Controllers: &Controllers{
			ListProduct: &listProductController,
			SaveProduct: &saveProductController,
		},
	}
}
