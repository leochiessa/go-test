package entities

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Uuid  uuid.UUID
	Name  string
	Price decimal.Decimal
}
