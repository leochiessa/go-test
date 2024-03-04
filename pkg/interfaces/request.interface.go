package interfaces

import (
	"context"
	"go-test/pkg/domain/dtos"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type RequestInterface interface {
	Connect()
	SelectAll(context.Context) pgx.Rows
	Select(context.Context, uuid.UUID) *dtos.RequestDto2
	Insert(context.Context, map[string]interface{})
	Update(context.Context, uuid.UUID, map[string]interface{}) int
	Delete(context.Context, uuid.UUID) int
}
