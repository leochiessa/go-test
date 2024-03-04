package services

import (
	"context"
	"go-test/pkg/domain/entities"
	"go-test/pkg/persistence/repositories"

	"github.com/fatih/structs"
	"github.com/gofrs/uuid"
)

type ClientService struct {
	Repository repositories.ClientRepository
}

func (s ClientService) GetAll(ctx context.Context) []entities.Client {
	rows := s.Repository.ClientInterface.SelectAll(ctx)
	var client entities.Client
	var clientsSlice []entities.Client
	for rows.Next() {
		rows.Scan(&client.Uuid, &client.Name, &client.Address)
		clientsSlice = append(clientsSlice, client)
	}
	return clientsSlice
}

func (s ClientService) Get(ctx context.Context, uuid uuid.UUID) entities.Client {
	rows := s.Repository.ClientInterface.Select(ctx, uuid)
	var client entities.Client
	for rows.Next() {
		rows.Scan(&client.Uuid, &client.Name, &client.Address)
	}
	return client
}

func (s ClientService) Save(ctx context.Context, c entities.Client) {
	c.Uuid, _ = uuid.NewV4()
	m := structs.Map(c)
	s.Repository.ClientInterface.Insert(ctx, m)
}

func (s ClientService) Edit(ctx context.Context, id uuid.UUID, c entities.Client) int {
	m := structs.Map(c)
	return s.Repository.ClientInterface.Update(ctx, id, m)
}

func (s ClientService) Remove(ctx context.Context, id uuid.UUID) int {
	return s.Repository.ClientInterface.Delete(ctx, id)
}
