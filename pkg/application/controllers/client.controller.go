package controllers

import (
	"encoding/json"
	"go-test/pkg/domain/entities"
	"go-test/pkg/domain/services"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)

type ClientController struct {
	Service services.ClientService
}

func (c ClientController) GetAll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.GetAll(ctx)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

func (c ClientController) Get(rw http.ResponseWriter, r *http.Request) {
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

func (c ClientController) Save(rw http.ResponseWriter, r *http.Request) {
	var client entities.Client
	json.NewDecoder(r.Body).Decode(&client)
	ctx := r.Context()
	c.Service.Save(ctx, client)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func (c ClientController) Edit(rw http.ResponseWriter, r *http.Request) {
	var client entities.Client
	json.NewDecoder(r.Body).Decode(&client)
	ctx := r.Context()
	res := c.Service.Edit(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")), client)
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (c ClientController) Remove(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := c.Service.Remove(ctx, uuid.FromStringOrNil(chi.URLParam(r, "uuid")))
	if res == 200 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
	} else if res == 409 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusConflict)
		json.NewEncoder(rw).Encode("This client cannot be deleted because it has one or more associated requests.")
	} else if res == 404 {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
	}
}
