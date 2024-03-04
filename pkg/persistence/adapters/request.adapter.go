package adapters

import (
	"context"
	"go-test/pkg/domain/dtos"
	"go-test/pkg/persistence/config"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RequestAdapter struct {
	Config         config.DataBaseConfig
	ConnectionPool *pgxpool.Pool
}

func (ra *RequestAdapter) Connect() {
	dbpool, err := pgxpool.New(context.Background(), ra.Config.GetConnectionString())
	ra.ConnectionPool = dbpool
	if err != nil {
		log.Fatal("Unable to connect to data base: ", err)
	}
}

func (ra *RequestAdapter) SelectAll(ctx context.Context) pgx.Rows {
	query, _, err := sq.Select("*").From("request").ToSql()
	rows, err := ra.ConnectionPool.Query(context.Background(), query)
	if err != nil {
		log.Fatal("Unable to select: ", err)
	}
	return rows
}

func (ra *RequestAdapter) Select(ctx context.Context, id uuid.UUID) *dtos.RequestDto2 {
	query, args, err := sq.Select("*").From("request").Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	rows, err := ra.ConnectionPool.Query(context.Background(), query, args...)
	var request dtos.RequestDto2
	for rows.Next() {
		rows.Scan(&request.Uuid, &request.Client.Uuid, &request.TimeStamp, &request.Status, &request.Total)
	}

	query, args, err = sq.Select("product_id, amount").From("request_detail").Where(sq.Eq{"request_id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	rows, err = ra.ConnectionPool.Query(context.Background(), query, args...)
	for rows.Next() {
		var requestDetail dtos.RequestDetailDto
		rows.Scan(&requestDetail.Product.Uuid, &requestDetail.Amount)
		request.RequestDetail = append(request.RequestDetail, requestDetail)
	}
	if err != nil {
		log.Fatal("Unable to select: ", err)
	}
	return &request
}

func (ra *RequestAdapter) Insert(ctx context.Context, entity map[string]interface{}) {
	values := make([]interface{}, 0, 5)
	values = append(values, entity["Uuid"])
	values = append(values, entity["Client"].(map[string]interface{})["Uuid"])
	values = append(values, "now()")
	values = append(values, entity["Status"])
	values = append(values, entity["Total"])

	query, _, _ := sq.Insert("request").Values(values...).PlaceholderFormat(sq.Dollar).ToSql()
	_, err := ra.ConnectionPool.Exec(context.Background(), query, values...)

	requestDetail := entity["RequestDetail"].([]interface{})
	for i := range requestDetail {
		requestDetailValues := make([]interface{}, 0, 3)
		requestDetailValues = append(requestDetailValues, entity["Uuid"])
		requestDetailValues = append(requestDetailValues, requestDetail[i].(map[string]interface{})["Product"].(map[string]interface{})["Uuid"])
		requestDetailValues = append(requestDetailValues, requestDetail[i].(map[string]interface{})["Amount"])
		query, _, _ := sq.Insert("request_detail").Values(requestDetailValues...).PlaceholderFormat(sq.Dollar).ToSql()
		_, err = ra.ConnectionPool.Exec(context.Background(), query, requestDetailValues...)
	}
	if err != nil {
		log.Fatal("Unable to insert: ", err)
	}
}

func (ra *RequestAdapter) Update(ctx context.Context, id uuid.UUID, entity map[string]interface{}) int {
	setMap := map[string]interface{}{
		"client_id":  entity["Client"].(map[string]interface{})["Uuid"],
		"time_stamp": "now()",
		"status":     entity["Status"],
		"total":      entity["Total"],
	}
	query, args, _ := sq.Update("request").SetMap(setMap).Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	ct, err1 := ra.ConnectionPool.Exec(context.Background(), query, args...)
	if ct.RowsAffected() != 0 {
		query, args, _ := sq.Delete("request_detail").Where(sq.Eq{"request_id": id}).PlaceholderFormat(sq.Dollar).ToSql()
		_, err2 := ra.ConnectionPool.Exec(context.Background(), query, args...)
		requestDetail := entity["RequestDetail"].([]interface{})
		for i := range requestDetail {
			requestDetailValues := make([]interface{}, 0, 3)
			requestDetailValues = append(requestDetailValues, id)
			requestDetailValues = append(requestDetailValues, requestDetail[i].(map[string]interface{})["Product"].(map[string]interface{})["Uuid"])
			requestDetailValues = append(requestDetailValues, requestDetail[i].(map[string]interface{})["Amount"])
			query, _, _ := sq.Insert("request_detail").Values(requestDetailValues...).PlaceholderFormat(sq.Dollar).ToSql()
			_, err2 = ra.ConnectionPool.Exec(context.Background(), query, requestDetailValues...)
		}
		if err2 != nil {
			log.Fatal("Unable to update: ", err2)
		}
		return 200
	}
	if err1 != nil {
		log.Fatal("Unable to update: ", err1)
	}
	return 404
}

func (ra *RequestAdapter) Delete(ctx context.Context, id uuid.UUID) int {
	query, args, _ := sq.Delete("request").Where(sq.Eq{"uuid": id}).PlaceholderFormat(sq.Dollar).ToSql()
	ct, err := ra.ConnectionPool.Exec(context.Background(), query, args...)
	if err != nil {
		log.Fatal("Unable to delete: ", err)
	}
	if ct.RowsAffected() != 0 {
		return 200
	}
	return 404
}
