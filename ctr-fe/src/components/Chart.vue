<template>
  <div style="width: 100%; height: 100%" class="container" :id="id"></div>
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
    id: {
      required: true,
      type: String,
      default() {
        return 'container'
      }
    },
  },
  data() {
    return {
      chart: {},
    }
  },
  mounted() {
    this.chart = echarts.init(document.getElementById(this.id), "light")
    this.chart.setOption(this.$props.options)
    this.chart.resize()
  },
  watch: {
    options: {
      handler(newOptions) {
        this.chart.clear()
        this.chart.resize()
        setTimeout(this.chart.setOption(newOptions), 500)
      },
      deep: true
    }
  },
  methods: {
    resize() {
      this.chart.resize()
    },
  }
}
</script>

<style scoped>
#container {
  width: 100%;
  height: 100%;
  background-image: url('../assets/logo.png');
  /*background-attachment: fixed;*/
  background-size: 200px 200px;
  background-repeat: no-repeat;
  background-position: center;
}
</style>