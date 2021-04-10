package model

import (
	"fmt"
	"log"
	"testing"
)

func TestMonitorInfluxDB(t *testing.T) {
	conn := connInflux()
	fmt.Println(conn)

	// insert
	writesPoints(conn)

	// 获取10条数据并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := queryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range res[0].Series[0].Values {
		for j, value := range row {
			log.Printf("j:%d value:%v\n", j, value)
		}
	}
}
