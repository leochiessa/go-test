package services

import (
	"context"
	"go-test/pkg/domain/entities"
	"go-test/pkg/persistence/repositories"

	"github.com/fatih/structs"
	"github.com/gofrs/uuid"
)

type ProductService struct {
	Repository repositories.ProductRepository
}

func (s ProductService) GetAll(ctx context.Context) []entities.Product {
	rows := s.Repository.ProductInterface.SelectAll(ctx)
	var product entities.Product
	var productsSlice []entities.Product
	for rows.Next() {
		rows.Scan(&product.Uuid, &product.Name, &product.Price)
		productsSlice = append(productsSlice, product)
	}
	return productsSlice
}

func (s ProductService) Get(ctx context.Context, uuid uuid.UUID) entities.Product {
	rows := s.Repository.ProductInterface.Select(ctx, uuid)
	var product entities.Product
	for rows.Next() {
		rows.Scan(&product.Uuid, &product.Name, &product.Price)
	}
	return product
}

func (s ProductService) Save(ctx context.Context, p entities.Product) {
	p.Uuid, _ = uuid.NewV4()
	m := structs.Map(p)
	s.Repository.ProductInterface.Insert(ctx, m)
}

func (s ProductService) Edit(ctx context.Context, id uuid.UUID, p entities.Product) int {
	m := structs.Map(p)
	return s.Repository.ProductInterface.Update(ctx, id, m)
}

func (s ProductService) Remove(ctx context.Context, id uuid.UUID) int {
	return s.Repository.ProductInterface.Delete(ctx, id)
}
