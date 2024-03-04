package controllers

import (
	"encoding/json"
	"go-test/pkg/domain/entities"
	"go-test/pkg/domain/services"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)

type RequestController struct {
	Service services.RequestService
}

func (c RequestController) GetAll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.GetAll(ctx)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

func (c RequestController) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.Get(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")))
	if res.Status != "" {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res)
	} else {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (c RequestController) Save(rw http.ResponseWriter, r *http.Request) {
	var request entities.Request
	json.NewDecoder(r.Body).Decode(&request)
	ctx := r.Context()
	c.Service.Save(ctx, request)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func (c RequestController) Edit(rw http.ResponseWriter, r *http.Request) {
	var request entities.Request
	json.NewDecoder(r.Body).Decode(&request)
	ctx := r.Context()
	res := c.Service.Edit(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")), request)
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (c RequestController) Remove(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.Remove(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")))
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}
