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
    this.chart = echarts.init(document.getElementById('container'), "light", {useDirtyRect: true})
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
  background-image: url('../assets/logo.png');
  /*background-attachment: fixed;*/
  background-size: 200px 200px;
  background-repeat: no-repeat;
  background-position: center;
}
</style>