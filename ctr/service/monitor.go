package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ahojcn/ecloud/ctr/model"
	client "github.com/influxdata/influxdb1-client/v2"
)

func getTableNameByMetrics(metrics string) string {
	return "host_" + metrics
}

func MonitorMetricsWrite(hostId int64, metrics string, data map[string]interface{}) {
	cli := model.GetInfluxDB()
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  "ecloud_monitor",
	})

	tags := map[string]string{"host_id": strconv.FormatInt(hostId, 10)}
	fields := data

	pt, _ := client.NewPoint(getTableNameByMetrics(metrics), tags, fields, time.Now())

	bp.AddPoint(pt)
	_ = cli.Write(bp)
}

func MonitorMetricsQuery(hostId int64, metrics string, cols []string, fromTime, toTime string) (res []client.Result, err error) {
	cli := model.GetInfluxDB()
	ss := []string{}
	for _, s := range cols {
		ss = append(ss, fmt.Sprintf("\"%s\"", s))
	}
	fmt.Println(ss, cols)
	cmd := fmt.Sprintf(`select %s from %s where host_id='%s' and time >= '%s' and time <= '%s'`,
		strings.Join(ss, ","),
		getTableNameByMetrics(metrics),
		strconv.FormatInt(hostId, 10),
		fromTime, toTime,
	)
	fmt.Println(cmd)

	q := client.Query{
		Command:  cmd,
		Database: "ecloud_monitor",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return nil, response.Error()
		}
		res = response.Results
	} else {
		return nil, err
	}
	return res, nil
}
