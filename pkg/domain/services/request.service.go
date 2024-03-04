package services

import (
	"context"
	"go-test/pkg/domain/dtos"
	"go-test/pkg/domain/entities"
	"go-test/pkg/persistence/repositories"

	"github.com/fatih/structs"
	"github.com/gofrs/uuid"
)

type RequestService struct {
	Repository repositories.RequestRepository
}

func (s RequestService) GetAll(ctx context.Context) []dtos.RequestDto1 {
	rows := s.Repository.RequestInterface.SelectAll(ctx)
	var request dtos.RequestDto1
	var requestsSlice []dtos.RequestDto1
	for rows.Next() {
		rows.Scan(&request.Uuid, &request.Client.Uuid, &request.TimeStamp, &request.Status, &request.Total)
		requestsSlice = append(requestsSlice, request)
	}
	return requestsSlice
}

func (s RequestService) Get(ctx context.Context, uuid uuid.UUID) *dtos.RequestDto2 {
	return s.Repository.RequestInterface.Select(ctx, uuid)
}

func (s RequestService) Save(ctx context.Context, r entities.Request) {
	r.Uuid, _ = uuid.NewV4()
	m := structs.Map(r)
	s.Repository.RequestInterface.Insert(ctx, m)
}

func (s RequestService) Edit(ctx context.Context, uuid uuid.UUID, r entities.Request) int {
	m := structs.Map(r)
	return s.Repository.RequestInterface.Update(ctx, uuid, m)
}

func (s RequestService) Remove(ctx context.Context, uuid uuid.UUID) int {
	return s.Repository.RequestInterface.Delete(ctx, uuid)
}
