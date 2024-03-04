package controllers

import (
	"encoding/json"
	"go-test/pkg/domain/entities"
	"go-test/pkg/domain/services"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)

type ProductController struct {
	Service services.ProductService
}

func (c ProductController) GetAll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.GetAll(ctx)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

func (c ProductController) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.Get(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")))
	if res.Name != "" {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res)
	} else {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (c ProductController) Save(rw http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	ctx := r.Context()
	c.Service.Save(ctx, product)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func (c ProductController) Edit(rw http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	ctx := r.Context()
	res := c.Service.Edit(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")), product)
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (c ProductController) Remove(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.Remove(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")))
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 409 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusConflict)
		json.NewEncoder(rw).Encode("This product cannot be deleted because it has one or more associated requests.")
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}
