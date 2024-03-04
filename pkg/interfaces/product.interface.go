package interfaces

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductInterface interface {
	Connect()
	SelectAll(context.Context) pgx.Rows
	Select(context.Context, uuid.UUID) pgx.Rows
	Insert(context.Context, map[string]interface{})
	Update(context.Context, uuid.UUID, map[string]interface{}) int
	Delete(context.Context, uuid.UUID) int
}
