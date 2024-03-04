package routers

import (
	"go-test/pkg/application/controllers"

	"github.com/go-chi/chi"
)

type RequestRouter struct {
	Router     *chi.Mux
	Controller controllers.RequestController
}

func (r RequestRouter) NewRequestRouter() *chi.Mux {
	r.Router = chi.NewRouter()
	r.Router.Get("/", r.Controller.GetAll)
	r.Router.Get("/{uuid}", r.Controller.Get)
	r.Router.Post("/", r.Controller.Save)
	r.Router.Patch("/{uuid}", r.Controller.Edit)
	r.Router.Delete("/{uuid}", r.Controller.Remove)
	return r.Router
}
