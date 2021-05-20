package service

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
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

func RouterMonitorMetricsGet() (map[string][]string, error) {
	cmd := "show series from router_logstash;"
	q := client.Query{Command: cmd, Database: "ecloud_monitor"}
	cli := model.GetInfluxDB()
	response, err := cli.Query(q)
	if err != nil {
		return nil, err
	}

	resp := map[string][]string{}
	for _, v := range response.Results[0].Series[0].Values {
		ss := strings.Split(v[0].(string), ",")
		k := strings.Split(ss[1], "=")[1]
		vv := strings.Split(ss[2], "=")[1]
		resp[k] = append(resp[k], vv)
	}
	return resp, nil
}

func RouterMonitorMetricsWrite(un string, uri string, data map[string]interface{}) {
	cli := model.GetInfluxDB()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  "ecloud_monitor",
	})
	if err != nil {
		fmt.Printf("influx db write failed: %v\n", err)
	}

	tags := map[string]string{"un": un, "uri": uri}
	fields := data
	pt, _ := client.NewPoint("router_logstash", tags, fields, time.Now())
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		fmt.Printf("influx db write failed: %v\n", err)
	}
}

func RouterMonitorMetricsQuery(rd *entity.RouterMonitorMetricsQueryRequestData) (res []client.Result, err error) {
	cli := model.GetInfluxDB()

	cmd := getRouterMonitorMetricsQueryCmd(rd)
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

func getRouterMonitorMetricsQueryCmd(rd *entity.RouterMonitorMetricsQueryRequestData) (cmd string) {
	cmd = fmt.Sprintf(
		"select %s from router_logstash where uri='%s' and un='%s' and time >= '%s' and time <= '%s'",
		*rd.Metrics,
		*rd.Uri, *rd.Un,
		*rd.FromTime, *rd.ToTime,
	)
	return
}


