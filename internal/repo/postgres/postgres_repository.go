package postgres

import (
	"database/sql"
	"encoding/json"
	"github.com/doug-martin/goqu/v9"
	"wb_l0/internal/repo"
)

type Repository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

var columnsToSelect = []interface{}{
	torder + ".order_uid",
	torder + ".track_number",
	torder + ".entry",
	torder + ".items",
	torder + ".locale",
	torder + ".internal_signature",
	torder + ".customer_id",
	torder + ".delivery_service",
	torder + ".shardkey",
	torder + ".sm_id",
	torder + ".date_created",
	torder + ".oof_shard",

	tdelivery + ".name",
	tdelivery + ".phone",
	tdelivery + ".zip",
	tdelivery + ".city",
	tdelivery + ".address",
	tdelivery + ".region",
	tdelivery + ".email",

	tpayment + ".transaction",
	tpayment + ".request_id",
	tpayment + ".currency",
	tpayment + ".provider",
	tpayment + ".amount",
	tpayment + ".payment_dt",
	tpayment + ".bank",
	tpayment + ".delivery_cost",
	tpayment + ".goods_total",
	tpayment + ".custom_fee",
}

func (r *Repository) All() ([]repo.OrderData, error) {
	var allRecords []repo.OrderData

	query, _, _ := goqu.From(torder).Select(columnsToSelect...).Join(
		goqu.T(tdelivery),
		goqu.On(goqu.Ex{torder + ".delivery_id": goqu.I(tdelivery + ".id")}),
	).Join(
		goqu.T(tpayment),
		goqu.On(goqu.Ex{torder + ".payment_id": goqu.I(tpayment + ".transaction")}),
	).ToSQL()

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		od := repo.OrderData{}

		var items string
		itemsToScan := []interface{}{
			&od.OrderUid,
			&od.TrackNumber,
			&od.Entry,
			&items,
			&od.Locale,
			&od.InternalSignature,
			&od.CustomerId,
			&od.DeliveryService,
			&od.ShardKey,
			&od.SmId,
			&od.DateCreated,
			&od.OofShard,

			&od.Delivery.Name,
			&od.Delivery.Phone,
			&od.Delivery.Zip,
			&od.Delivery.City,
			&od.Delivery.Address,
			&od.Delivery.Region,
			&od.Delivery.Email,

			&od.Payment.Transaction,
			&od.Payment.RequestId,
			&od.Payment.Currency,
			&od.Payment.Provider,
			&od.Payment.Amount,
			&od.Payment.PaymentDt,
			&od.Payment.Bank,
			&od.Payment.DeliveryCost,
			&od.Payment.GoodsTotal,
			&od.Payment.CustomFee,
		}

		err := rows.Scan(itemsToScan...)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal([]byte(items), &od.Items); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, od)
	}
	return allRecords, nil
}

func (r *Repository) GetById(id string) (*repo.OrderData, error) {
	query, _, _ := goqu.From(torder).Select(columnsToSelect...).Join(
		goqu.T(tdelivery),
		goqu.On(goqu.Ex{torder + ".delivery_id": goqu.I(tdelivery + ".id")}),
	).Where(goqu.Ex{torder + ".order_uid": goqu.I(id)}).Join(
		goqu.T(tpayment),
		goqu.On(goqu.Ex{torder + ".payment_id": goqu.I(tpayment + ".transaction")}),
	).ToSQL()

	row := r.db.QueryRow(query)

	od := repo.OrderData{}

	var items string
	itemsToScan := []interface{}{
		&od.OrderUid,
		&od.TrackNumber,
		&od.Entry,
		&items,
		&od.Locale,
		&od.InternalSignature,
		&od.CustomerId,
		&od.DeliveryService,
		&od.ShardKey,
		&od.SmId,
		&od.DateCreated,
		&od.OofShard,

		&od.Delivery.Name,
		&od.Delivery.Phone,
		&od.Delivery.Zip,
		&od.Delivery.City,
		&od.Delivery.Address,
		&od.Delivery.Region,
		&od.Delivery.Email,

		&od.Payment.Transaction,
		&od.Payment.RequestId,
		&od.Payment.Currency,
		&od.Payment.Provider,
		&od.Payment.Amount,
		&od.Payment.PaymentDt,
		&od.Payment.Bank,
		&od.Payment.DeliveryCost,
		&od.Payment.GoodsTotal,
		&od.Payment.CustomFee,
	}

	err := row.Scan(itemsToScan...)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(items), &od.Items); err != nil {
		return nil, err
	}

	return &od, nil
}

func (r *Repository) Insert(od repo.OrderData) error {
	insert := goqu.Insert(tdelivery).Rows(od.Delivery)
	query, _, _ := insert.ToSQL()
	var deliveryId int
	err := r.db.QueryRow(query + " RETURNING id").Scan(&deliveryId)
	if err != nil {
		return err
	}

	insert = goqu.Insert(tpayment).Rows(od.Payment)
	query, _, _ = insert.ToSQL()
	if _, err := r.db.Exec(query); err != nil {
		return err
	}

	items, _ := json.Marshal(od.Items)
	insert = goqu.Insert(torder).Rows(
		goqu.Record{
			"order_uid":          od.OrderUid,
			"track_number":       od.TrackNumber,
			"entry":              od.Entry,
			"delivery_id":        deliveryId,
			"payment_id":         od.Payment.Transaction,
			"items":              items,
			"locale":             od.Locale,
			"internal_signature": od.InternalSignature,
			"customer_id":        od.CustomerId,
			"delivery_service":   od.DeliveryService,
			"shardkey":           od.ShardKey,
			"sm_id":              od.SmId,
			"date_created":       od.DateCreated,
			"oof_shard":          od.OofShard,
		},
	)
	query, _, _ = insert.ToSQL()
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
