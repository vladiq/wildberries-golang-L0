package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"net/http"
	"time"
	"wb_l0/internal"
	"wb_l0/internal/api"
	"wb_l0/internal/repo"
	"wb_l0/internal/repo/postgres"
	"wb_l0/internal/services/data_generator"
)

import (
	_ "github.com/lib/pq"
	cache2 "wb_l0/internal/repo/cache"
)

const (
	pgConnInfo          = "host=127.0.0.1 port=5432 user=vladislav password=1 dbname=wb_l0 sslmode=disable"
	pgDriver            = "postgres"
	stanClusterID       = "test-cluster"
	stanURL             = "0.0.0.0:4222"
	stanClientID        = "wb-l0-subscriber"
	subscriptionSubject = "order_data"
)

func main() {
	sc, err := stan.Connect(stanClusterID, stanClientID, stan.NatsURL(stanURL))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	db, err := sql.Open(pgDriver, pgConnInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgData := repo.Repository(postgres.NewPostgresRepository(db))
	cacheData := repo.Repository(cache2.NewCacheRepository())
	if err := fillCacheFromDB(pgData, cacheData); err != nil {
		panic(err)
	}

	sub, _ := sc.Subscribe(subscriptionSubject, func(m *stan.Msg) {
		if err := processMessage(m, cacheData, pgData); err != nil {
			panic(err)
		}
	})
	defer sub.Unsubscribe()

	server := api.NewServer(&cacheData)
	if err := http.ListenAndServe("localhost:8080", server); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Minute)
}

func fillCacheFromDB(pgData repo.Repository, cacheData repo.Repository) error {
	if recordsFromDb, err := pgData.All(); err != nil {
		return err
	} else {
		for _, record := range recordsFromDb {
			cacheData.Insert(record)
		}
	}
	return nil
}

func processMessage(m *stan.Msg, cacheData repo.Repository, pgData repo.Repository) error {
	parsedJson := data_generator.ModelJSON{}
	if err := json.Unmarshal(m.Data, &parsedJson); err != nil {
		errorMsg := fmt.Errorf("unmarshalling struct: %w. Skipping", err)
		fmt.Println(errorMsg)
	} else {
		orderData := *internal.MapGeneratedToStored(&parsedJson)
		cacheData.Insert(orderData)
		if err := pgData.Insert(orderData); err != nil {
			return err
		}
	}
	return nil
}
