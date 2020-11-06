package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "coins"
	username = ""
	password = ""
)

func main() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"name": "test"}

	coins := map[string]interface{}{"coin": 123.123}
	pt, err := client.NewPoint("myuser1", tags, coins, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := conn.Write(bp); err != nil {
		log.Fatal(err)
	}
}
