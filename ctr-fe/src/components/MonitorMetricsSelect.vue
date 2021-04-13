<template>
  <div style="width: 100%; height: 100%">
    <Tabs @on-click="onClickTab">
      <TabPane v-for="(metric,index) in metrics" :key="index"
               :label="metric" :name="metric">
        <CheckboxGroup v-model="cols_selected">
          <Checkbox v-for="(col, index) in cols[metric]" :key="index" :label="col" border></Checkbox>
        </CheckboxGroup>
      </TabPane>
    </Tabs>

    <div>
      <Button type="primary" @click="onClickBtnGet">GET！</Button>
    </div>
  </div>
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
          "load1", "load5", "load15", "process_count"
        ],
        "net": [
          "bytes_sent", "bytes_recv", "packets_sent", "packets_recv", "err_in", "err_out", "drop_in",
          "drop_out", "fifo_in", "fifo_out"
        ]
      },
      cols_selected: [],
    }
  },
  methods: {
    onClickTab(name) {
      this.metrics_selected = name
    },
    onClickBtnGet() {
      if (this.metrics_selected !== '' && this.cols_selected.length !== 0) {
        this.$emit('onMonitorMetricsSelected', this.metrics_selected, this.cols_selected)
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