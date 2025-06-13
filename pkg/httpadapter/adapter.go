package httpadapter

import "net/http"

type IHttpAdapter interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}
