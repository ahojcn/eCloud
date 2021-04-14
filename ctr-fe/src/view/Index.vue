<template>
  <div>
    <Menu mode="horizontal" theme="light" active-name="1" @on-select="handleMenuSelect">
      <div>
        <MenuItem name="service">
          <Icon type="md-infinite"/>
          服务
        </MenuItem>
        <MenuItem name="perm">
          <Icon type="ios-people"/>
          权限
        </MenuItem>
        <MenuItem name="monitor">
          <Icon type="ios-stats"/>
          监控
        </MenuItem>
        <MenuItem name="deploy">
          <Icon type="md-cloud-upload"/>
          部署
        </MenuItem>
        <MenuItem name="resource">
          <Icon type="md-code-working"/>
          资源
        </MenuItem>
        <MenuItem name="dev">
          <Icon type="ios-bug"/>
          开发
        </MenuItem>
      </div>

      <div style="margin-right: 12px;float: right">
        <Submenu name="avatar">
          <template slot="title">
            <Avatar size="large" style="color: #f56a00;background-color: #fde3cf">{{ user_info.username }}</Avatar>
          </template>
          <MenuItem name="logout">
            <Icon type="md-exit"/>
            退出
          </MenuItem>
        </Submenu>
      </div>
    </Menu>

    <MonitorMetricsSelect @onMonitorMetricsSelected="onMonitorMetricsSelected"></MonitorMetricsSelect>
    <Chart :options="options" style="width: 1000px; height: 500px"></Chart>
    <ServiceTree @onTreeNodeSelected="onTreeNodeSelected"></ServiceTree>
  </div>
</template>

<script>
import {is_login, logout} from "@/api/session";
import ServiceTree from "@/components/ServiceTree";
import Chart from "@/components/Chart";
import {get_metrics_data} from "@/api/monitor"
import MonitorMetricsSelect from "@/components/MonitorMetricsSelect";

export default {
  name: "Index",
  components: {MonitorMetricsSelect, Chart, ServiceTree},
  data() {
    return {
      options: {
        title: {
          text: '监控图'
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
          data: ['user', 'system']
        },
        xAxis: {
          type: 'time',
        },
        yAxis: {
          type: 'value',
        },
        series: []
      }
    }
  },
  computed: {
    user_info() {
      return this.$store.state.user_info
    }
  },
  mounted() {
    is_login().then(res => {
      if (res.code !== 200) {
        this.$router.push("/login")
      }
      this.$store.commit('set_user_info', res.data)
    })
  },
  methods: {
    handleMenuSelect(name) {
      if (name === 'logout') {
        this.handleClickLogout()
      }
    },
    handleClickLogout() {
      logout().then(res => {
        if (res.code === 200) {
          this.$router.go(0)
        }
      })
    },
    onTreeNodeSelected(node) {
      console.log(node)
    },
    onMonitorMetricsSelected(metric, cols) {
      this.options.legend.data = cols
      this.options.series.length = 0
      this.options.title.text = metric
      for (let i = 0; i < cols.length; i++) {
        get_metrics_data({
          "host_id": '12',
          "metrics": metric,
          "cols": cols[i]
        }).then(res => {
          this.options.series.push({type: "line", name: cols[i], data: res.data[0].Series[0].values})
        })
      }
    }
  }
}
</script>

<style scoped>
</style>