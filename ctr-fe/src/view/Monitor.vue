<template>
  <div>
    <Row :gutter="18">
      <Col :span="6">
        <Card title="指标选择" style="min-height: 80vh">
          <div slot="extra">
            <Button type="primary" @click="onRouterMonitorMetricsBtnClick">GET！</Button>
          </div>
          <Form>
            <FormItem label="一级指标">
              <Cascader :data="metrics" v-model="selected" filterable></Cascader>
              <span style="color: red">{{ selected }}</span>
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
              <br>
              <span style="color: red">{{ datetime }}</span>
            </FormItem>
          </Form>
          <Chart ref="overview_chart" style="height: 300px" :options="overview_chart_options" id="overview"></Chart>
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
      selected: ['test-1.testsvc.frontend.xiaoniu.zhieasy', '/'],
      query: [
        {value: 'status', label: 'status(状态码)'},
        {value: 'request_time', label: 'request_time(请求->响应时间)'},
        {value: 'upstream_response_time', label: 'upstream_response_time(请求->分发请求)'},
      ],
      query_selected: ['request_time', 'upstream_response_time'],
      datetime: [new Date('2021-05-19T10:32:00+08:00'), new Date('2021-05-21T10:33:00+08:00')],

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
      overview_chart_options_raw: {
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
        ]
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
        ]
      },
    }
  },
  computed: {
    metrics: function () {
      let data = []
      for (let k in this.raw_resp) {
        let children = []

        for (let i = 0; i < this.raw_resp[k].length; i++) {
          let uri = this.raw_resp[k][i]
          children.push({value: uri, label: `uri=${uri}`})
        }

        data.push({value: k, label: `un=${k}`, children: children})
      }
      return data
    }
  },
  mounted() {
    apiGetRouterMetrics().then(res => {
      if (res.code === 200) {
        this.raw_resp = res.data
      }
    })
  },
  methods: {
    onRouterMonitorMetricsBtnClick() {
      this.chart_options = this.chart_options_raw
      this.chart_options.series.length = 0
      this.chart_options.title.text = JSON.stringify(this.query_selected)
      for (let i = 0; i < this.query_selected.length; i++) {
        apiQueryRouterMetrics({
          un: this.selected[0],
          uri: this.selected[1],
          metrics: this.query_selected[i],
          from_time: this.datetime[0],
          to_time: this.datetime[1],
          overview: false,
        }).then(res => {
          this.chart_options.series.push({
            type: "line",
            name: this.query_selected[i],
            data: res.data[0].Series[0].values
          })
          this.chart_options.legend.data.push(this.query_selected[i])
        })
      }

      // 获取 overview 数据
      apiQueryRouterMetrics({
        un: this.selected[0],
        uri: this.selected[1],
        metrics: "",
        from_time: this.datetime[0],
        to_time: this.datetime[1],
        overview: true,
      }).then(res => {
        this.overview_chart_options.series[0].data = res.data
      });
    },
  },
}
</script>

<style scoped>

</style>