package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
	"wb_l0/internal/services/data_generator"
)

const (
	stanClusterID = "test-cluster"
	stanClientID  = "wb-l0-wb-l0-publisher"
	stanURL       = "0.0.0.0:4222"
)

func main() {
	sc, err := stan.Connect(stanClusterID, stanClientID, stan.NatsURL(stanURL))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sc.Close()

	for i := 0; i < 100; i++ {
		dataToPublish := data_generator.Generate()
		publishSubject := "order_data"
		if err := sc.Publish(publishSubject, dataToPublish); err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}
