package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type AppRouter struct {
	Router        *chi.Mux
	ClientRouter  ClientRouter
	ProductRouter ProductRouter
	RequestRouter RequestRouter
}

func (r *AppRouter) NewAppRouter() *chi.Mux {
	r.Router = chi.NewRouter()
	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.StripSlashes)
	r.Router.Mount("/client", r.ClientRouter.Router)
	r.Router.Mount("/product", r.ProductRouter.Router)
	r.Router.Mount("/request", r.RequestRouter.Router)
	return r.Router
}
