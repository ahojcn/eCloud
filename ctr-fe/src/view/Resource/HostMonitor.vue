<template>
  <div style="height: 70vh">
    <MonitorMetricsSelect @onMonitorMetricsSelected="onMonitorMetricsSelected"></MonitorMetricsSelect>
    <Chart :options="monitor_options" id="host_monitor"></Chart>
  </div>
</template>

<script>
import {apiGetMetricsData} from "@/api/monitor";

import Chart from "@/components/Chart";
import MonitorMetricsSelect from "@/components/MonitorMetricsSelect";

export default {
  name: "HostMonitor",
  components: {MonitorMetricsSelect, Chart},
  data() {
    return {
      monitor_options: {
        title: {
          text: '监控图'
        },
        animation: false,
        dataZoom: [
          {
            id: 'dataZoomX',
            type: 'slider',
            xAxisIndex: [0],
            filterMode: 'filter', // 设定为 'filter' 从而 X 的窗口变化会影响 Y 的范围。
            start: 0,
            end: 100
          },
          {
            id: 'dataZoomY',
            type: 'slider',
            yAxisIndex: [0],
            filterMode: 'filter',
            start: 0,
            end: 100
          }
        ],
        toolbox: {
          show: true,
          showTitle: false, // 隐藏默认文字，否则两者位置会重叠
          feature: {
            saveAsImage: {
              show: true,
              title: 'Save As Image'
            },
            magicType: {
              type: ['line', 'bar', 'stack', 'tiled']
            },
            dataView: {
              show: true,
              title: 'Data View'
            },
            dataZoom: {
              show: true,
              title: '缩放'
            }
          },
        },
        tooltip: {
          show: true,
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
          },
        },
        legend: {
          data: []
        },
        xAxis: {
          type: 'time',
          splitLine: {
            show: false
          },
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, '100%'],
          splitLine: {
            show: false
          }
        },
        series: [],
      },
    }
  },
  computed: {
    id: {
      get() {
        return this.$route.query.id
      },
    }
  },
  methods: {
    onMonitorMetricsSelected(metrics, cols, from_time, to_time) {
      this.monitor_options.series.length = 0
      this.monitor_options.title.text = metrics + '(' + cols + ')'
      for (let i = 0; i < cols.length; i++) {
        apiGetMetricsData({
          "host_id": this.id,
          "metrics": metrics,
          "cols": cols[i],
          "from_time": from_time,
          "to_time": to_time
        }).then(res => {
          this.monitor_options.series.push({
            type: "line", name: cols[i], data: res.data[0].Series[0].values,
            areaStyle: {}, showSymbol: false, hoverAnimation: false,
          })
          this.monitor_options.legend.data.push(cols[i])
        })
      }
    },
  }
}
</script>

<style scoped>

</style>