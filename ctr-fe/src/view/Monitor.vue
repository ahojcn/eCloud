<template>
  <div>
    <Row :gutter="18">
      <Col :span="6">
        <Chart ref="overview_chart" style="height: 300px" :options="overview_chart_options" id="overview"></Chart>
        <Card title="指标选择" style="min-height: 50vh">
          <div slot="extra">
            <Tooltip :content="auto_refresh===true?'关闭自动刷新':'开启自动刷新'">
              <i-switch v-model="auto_refresh" @on-change="onAutoRefreshChange"></i-switch>
            </Tooltip>
            <Button type="primary" @click="onRouterMonitorMetricsBtnClick">GET！</Button>
          </div>
          <Form>
            <FormItem label="一级指标">
              <Cascader :data="metrics" v-model="selected" filterable></Cascader>
            </FormItem>
            <FormItem label="二级指标">
              <Select multiple v-model="query_selected">
                <Option v-for="(item, index) in query" :key="index" :value="item.value">
                  {{ item.label }}
                </Option>
              </Select>
            </FormItem>
            <FormItem lebal="时间区间">
              <DatePicker style="width: 100%" type="datetimerange" placeholder="选择时间段"
                          v-model="datetime"></DatePicker>
            </FormItem>
          </Form>
          <Alert type="error">
            当前已选
            <template slot="desc">
              <p>un：{{ selected[0] }}</p>
              <p>uri：{{ selected[1] }}</p>
              <p>docker：{{ selected[2] }}</p>
              <p>二级指标：{{ query_selected }}</p>
              <p>开始：{{ new Date(datetime[0]).toLocaleString() }}</p>
              <p>结束：{{ new Date(datetime[1]).toLocaleString() }}</p>
            </template>
          </Alert>
        </Card>
      </Col>
      <Col :span="18">
        <Chart ref="detail_chart" :options="chart_options" id="detail"></Chart>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiGetRouterMetrics, apiQueryRouterMetrics} from "@/api/router";
import Chart from "@/components/Chart";

export default {
  name: "Monitor",
  components: {Chart},
  data() {
    return {
      raw_resp: {},
      selected: ['test-1.testsvc.frontend.xiaoniu.zhieasy', '/', 'all'],
      query: [
        {value: 'status', label: 'status(状态码)'},
        {value: 'request_time', label: 'request_time(请求->响应时间)'},
        {value: 'upstream_response_time', label: 'upstream_response_time(请求->分发请求)'},
      ],
      query_selected: ['request_time', 'upstream_response_time'],
      datetime: [new Date('2021-05-19T10:32:00+08:00'), new Date('2021-05-21T10:33:00+08:00')],

      auto_refresh: false,
      auto_refresh_func: null,

      chart_options_raw: {
        title: {
          text: '监控图',
          left: 'center',
        },
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
          data: [],
          left: 10,
          orient: 'vertical',
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
      chart_options: {
        title: {
          text: '监控图',
          left: 'center',
        },
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
          data: [],
          left: 10,
          orient: 'vertical',
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
      overview_chart_options: {
        title: {
          text: '概览',
          left: 'center'
        },
        tooltip: {
          trigger: 'item'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
        },
        series: [
          {
            name: '请求状态码',
            type: 'pie',
            radius: '50%',
            data: [],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ],
      },
    }
  },
  computed: {
    metrics: function () {
      let data = []
      for (let k in this.raw_resp) {
        let children_uri = []
        let uri_map = this.raw_resp[k]
        for (let uri in uri_map) {
          let children_docker = []
          children_docker.push({value: 'all', label: `docker=*all*`})
          for (let i = 0; i < uri_map[uri].length; i++) {
            let docker = uri_map[uri][i]
            children_docker.push({value: docker, label: `docker=${docker}`})
          }
          children_uri.push({value: uri, label: `uri=${uri}`, children: children_docker})
        }
        data.push({value: k, label: `un=${k}`, children: children_uri})
      }
      return data
    }
  },
  mounted() {
    apiGetRouterMetrics().then(res => {
      if (res.code === 200) {
        this.raw_resp = res.data
        this.onRouterMonitorMetricsBtnClick()
      }
    })
  },
  methods: {
    onRouterMonitorMetricsBtnClick(animation) {
      if (this.selected.length !== 3 || this.query_selected.length === 0) {
        this.$Message.error('请按照要求选择一级指标和二级指标')
        return
      }

      this.chart_options.series = []
      this.chart_options.title.text = JSON.stringify(this.query_selected)
      for (let i = 0; i < this.query_selected.length; i++) {
        apiQueryRouterMetrics({
          un: this.selected[0],
          uri: this.selected[1],
          docker: this.selected[2],
          metrics: this.query_selected[i],
          from_time: this.datetime[0],
          to_time: this.datetime[1],
          overview: false,
        }).then(res => {
          if (res.code === 200) {
            this.chart_options.series.push({
              type: "line",
              showSymbol: false,
              name: this.query_selected[i],
              data: res.data[0].Series[0].values
            })
            this.chart_options.animation = animation
            this.chart_options.legend.data.push(this.query_selected[i])
          }
        })
      }

      // 获取 overview 数据
      apiQueryRouterMetrics({
        un: this.selected[0],
        uri: this.selected[1],
        docker: this.selected[2],
        metrics: "",
        from_time: this.datetime[0],
        to_time: this.datetime[1],
        overview: true,
      }).then(res => {
        if (res.code === 200) {
          this.overview_chart_options.animation = animation
          this.overview_chart_options.series[0].data = res.data
        }
      });
    },
    onAutoRefreshChange(status) {
      if (status === true) {
        this.auto_refresh_func = setInterval(() => {
          this.onRouterMonitorMetricsBtnClick(false)
        }, 2000)
        this.$Message.success('定时刷新已开启，2s')
      } else {
        clearInterval(this.auto_refresh_func)
        this.$Message.success('定时刷新已关闭')
      }
    },
  },
}
</script>

<style scoped>

</style>