package app

import (
	"fmt"

	"httpadapter"
)

var dependencies *Dependencies

type Api struct {
	deps *Dependencies
}

func New() *Api {
	api := Api{InitDependencies()}

	api.deps.Echo.GET("/product", httpadapter.EchoAdapter(api.deps.Controllers.ListProduct))
	api.deps.Echo.POST("/product", httpadapter.EchoAdapter(api.deps.Controllers.SaveProduct))

	return &api
}

func (api *Api) Start(address string, port int) {
	api.deps.Echo.Start(
		fmt.Sprintf("%s:%d", address, port),
	)
}
