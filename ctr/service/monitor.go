package service

import (
	"encoding/json"
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

func RouterMonitorMetricsGet() (map[string]map[string][]string, error) {
	cmd := "show series from router_logstash;"
	q := client.Query{Command: cmd, Database: "ecloud_monitor"}
	cli := model.GetInfluxDB()
	response, err := cli.Query(q)
	if err != nil {
		return nil, err
	}

	resp := map[string]map[string][]string{}
	uriMap := map[string][]string{}
	for _, v := range response.Results[0].Series[0].Values {
		ss := strings.Split(v[0].(string), ",")
		docker := strings.Split(ss[1], "=")[1]
		un := strings.Split(ss[2], "=")[1]
		uri := strings.Split(ss[3], "=")[1]
		uriMap[uri] = append(uriMap[uri], docker)
		resp[un] = uriMap
	}
	return resp, nil
}

func RouterMonitorMetricsWrite(un string, uri string, upstreamAddr string, data map[string]interface{}) {
	cli := model.GetInfluxDB()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  "ecloud_monitor",
	})
	if err != nil {
		fmt.Printf("influx db write failed: %v\n", err)
	}

	tags := map[string]string{"un": un, "uri": uri, "docker": upstreamAddr}
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

func RouterMonitorMetricsQueryOverview(rd *entity.RouterMonitorMetricsQueryRequestData) (res []map[string]interface{}, err error) {
	cli := model.GetInfluxDB()
	cmd := ""
	if *rd.Docker == "all" {
		cmd = fmt.Sprintf(
			"select distinct(status) from router_logstash where uri='%s' and un='%s' and time >= '%s' and time <= '%s';",
			*rd.Uri, *rd.Un,
			*rd.FromTime, *rd.ToTime,
		)
	} else {
		cmd = fmt.Sprintf(
			"select distinct(status) from router_logstash where uri='%s' and un='%s' and docker='%s' and time >= '%s' and time <= '%s';",
			*rd.Uri, *rd.Un, *rd.Docker,
			*rd.FromTime, *rd.ToTime,
		)
	}
	q := client.Query{Command: cmd, Database: "ecloud_monitor"}
	response, err := cli.Query(q)
	if err != nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}

	for _, row := range response.Results[0].Series[0].Values {
		if *rd.Docker == "all" {
			cmd = fmt.Sprintf(
				"select count(*) from router_logstash where uri='%s' and un='%s' and time >= '%s' and time <= '%s' and status=%v;",
				*rd.Uri, *rd.Un,
				*rd.FromTime, *rd.ToTime,
				row[1],
			)
		} else {
			cmd = fmt.Sprintf(
				"select count(*) from router_logstash where uri='%s' and un='%s' and docker='%s' and time >= '%s' and time <= '%s' and status=%v;",
				*rd.Uri, *rd.Un, *rd.Docker,
				*rd.FromTime, *rd.ToTime,
				row[1],
			)
		}
		q.Command = cmd
		rr, err := cli.Query(q)
		if err != nil {
			return nil, err
		}
		if rr.Error() != nil {
			return nil, rr.Error()
		}
		mm := map[string]interface{}{}
		mm["value"] = rr.Results[0].Series[0].Values[0][1]
		mm["name"] = row[1].(json.Number).String()
		res = append(res, mm)
	}

	return res, nil
}

func getRouterMonitorMetricsQueryCmd(rd *entity.RouterMonitorMetricsQueryRequestData) (cmd string) {
	if *rd.Docker == "all" {
		cmd = fmt.Sprintf(
			"select %s from router_logstash where uri='%s' and un='%s' and time >= '%s' and time <= '%s'",
			*rd.Metrics,
			*rd.Uri, *rd.Un,
			*rd.FromTime, *rd.ToTime,
		)
	} else {
		cmd = fmt.Sprintf(
			"select %s from router_logstash where uri='%s' and un='%s' and docker='%s' and time >= '%s' and time <= '%s'",
			*rd.Metrics,
			*rd.Uri, *rd.Un, *rd.Docker,
			*rd.FromTime, *rd.ToTime,
		)
	}
	return
}
