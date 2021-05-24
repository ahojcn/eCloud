<template>
  <Card>
    <div style="width: 100%; height: 100%">
      <Tabs @on-click="onClickTab">
        <template slot="extra">
          自动刷新
          <i-switch v-model="auto_refresh" @on-change="onAutoFreshChange" true-color="#13ce66" false-color="#ff4949" size="large"
                    slot="extra">
            <span slot="open">开</span>
            <span slot="close">关</span>
          </i-switch>
        </template>
        <TabPane v-for="(metric,index) in metrics" :key="index"
                 :label="metric" :name="metric">
          <CheckboxGroup v-model="cols_selected">
            <Checkbox v-for="(col, index) in cols[metric]" :key="index" :label="col" border></Checkbox>
          </CheckboxGroup>
        </TabPane>
      </Tabs>

      <div>
        <DatePicker type="datetimerange" placeholder="选择时间段"
                    style="width: 400px" v-model="datetime"></DatePicker>
        <Button type="primary" @click="onClickBtnGet">GET！</Button>
      </div>
    </div>
  </Card>
</template>

<script>
export default {
  name: "MonitorMetricsSelect",
  data() {
    return {
      metrics: [
        "cpu", "mem", "disk", "load", "net"
      ],
      metrics_selected: 'cpu',
      cols: {
        "cpu": [
          "user", "system", "idle", "percent"
        ],
        "mem": [
          "swap_total", "swap_used", "swap_free", "swap_percent", "virtual_total", "virtual_available", "virtual_used",
          "virtual_percent", "virtual_free", "virtual_active", "virtual_inactive", "virtual_wired", "virtual_buffers",
          "virtual_cached"
        ],
        "disk": [
          "total", "free", "used", "percent", "inodes_total", "inodes_used", "inodes_free", "inodes_percent"
        ],
        "load": [
          "load_1", "load_5", "load_15", "process_count"
        ],
        "net": [
          "bytes_sent", "bytes_recv", "packets_sent", "packets_recv", "err_in", "err_out", "drop_in",
          "drop_out", "fifo_in", "fifo_out"
        ]
      },
      cols_selected: ['percent'],
      datetime: [new Date('2021-05-10T10:32:00+08:00'), new Date('2021-05-29T10:33:00+08:00')],

      auto_refresh: false,
      auto_refresh_id: null,
    }
  },
  mounted() {
    this.$emit('onMonitorMetricsSelected', this.metrics_selected, this.cols_selected, this.datetime[0], this.datetime[1])
  },
  methods: {
    onClickTab(name) {
      this.metrics_selected = name
      this.cols_selected.length = 0
    },
    onAutoFreshChange(status) {
      if (status === true) {
        this.auto_refresh_id = setInterval(()=>{
          this.onClickBtnGet()
        }, 2000)
        this.$Message.success('自动刷新已开启：2s')
      } else {
        clearInterval(this.auto_refresh_id)
        this.$Message.info('自动刷新已关闭')
      }
    },
    onClickBtnGet() {
      if (this.metrics_selected !== '' && this.cols_selected.length !== 0) {
        this.$emit('onMonitorMetricsSelected', this.metrics_selected, this.cols_selected, this.datetime[0], this.datetime[1])
      } else {
        this.$Notice.warning({
          title: '请选一个指标'
        })
      }
    },
  }
}
</script>

<style scoped>

</style>