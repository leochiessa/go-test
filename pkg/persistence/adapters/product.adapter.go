package adapters

import (
	"context"
	"errors"
	"go-test/pkg/persistence/config"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductAdapter struct {
	Config         config.DataBaseConfig
	ConnectionPool *pgxpool.Pool
}

func (pa *ProductAdapter) Connect() {
	dbpool, err := pgxpool.New(context.Background(), pa.Config.GetConnectionString())
	pa.ConnectionPool = dbpool
	if err != nil {
		log.Fatal("Unable to connect to data base: ", err)
	}
}

func (pa *ProductAdapter) SelectAll(ctx context.Context) pgx.Rows {
	query, _, err := sq.Select("*").From("product").ToSql()
	rows, err := pa.ConnectionPool.Query(context.Background(), query)
	if err != nil {
		log.Fatal("Unable to select: ", err)
	}
	return rows
}

func (pa *ProductAdapter) Select(ctx context.Context, id uuid.UUID) pgx.Rows {
	query, args, err := sq.Select("*").From("product").Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	rows, err := pa.ConnectionPool.Query(context.Background(), query, args...)
	if err != nil {
		log.Fatal("Unable to select: ", err)
	}
	return rows
}

func (pa *ProductAdapter) Insert(ctx context.Context, entity map[string]interface{}) {
	values := make([]interface{}, 0, len(entity))
	for _, v := range entity {
		values = append(values, v)
	}
	query, _, _ := sq.Insert("product").Values(values...).PlaceholderFormat(sq.Dollar).ToSql()
	_, err := pa.ConnectionPool.Exec(context.Background(), query, values...)
	if err != nil {
		log.Fatal("Unable to insert: ", err)
	}
}

func (pa *ProductAdapter) Update(ctx context.Context, id uuid.UUID, entity map[string]interface{}) int {
	delete(entity, "Uuid")
	query, args, _ := sq.Update("product").SetMap(entity).Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	ct, err := pa.ConnectionPool.Exec(context.Background(), query, args...)
	if err != nil {
		log.Fatal("Unable to update: ", err)
	}
	if ct.RowsAffected() != 0 {
		return 200
	}
	return 404
}

func (pa *ProductAdapter) Delete(ctx context.Context, id uuid.UUID) int {
	query, args, _ := sq.Delete("product").Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	ct, err := pa.ConnectionPool.Exec(context.Background(), query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.SQLState() == "23503" {
			return 409
		}
		log.Fatal("Unable to delete: ", err)
	}
	if ct.RowsAffected() != 0 {
		return 200
	}
	return 404
}
