<template>
  <div class="container" id="container"></div>
</template>

<script>
import * as echarts from "echarts"

export default {
  name: "Chart",
  props: {
    options: {
      required: true,
      type: Object,
      default() {
        return {}
      },
    },
  },
  data() {
    return {
      chart: {},
    }
  },
  mounted() {
    this.chart = echarts.init(document.getElementById('container'))
    this.chart.setOption(this.$props.options)
  },
  watch: {
    options: {
      handler(newOptions) {
        this.chart.clear()
        setTimeout(this.chart.setOption(newOptions), 500)
      },
      deep: true
    }
  }
}
</script>

<style scoped>
#container {
  width: 100%;
  height: 100%;
}
</style>