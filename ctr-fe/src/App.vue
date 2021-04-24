<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      model: 'z16'
    }
  },
  created() {
    // window.L2Dwidget.init({
    //   pluginRootPath: '/static/live2dw/',
    //   pluginJsPath: 'lib/',
    //   pluginModelPath: 'live2d-widget-model-wanko/assets/',
    //   model: {jsonPath: `/static/live2dw/live2d-widget-model-${this.model}/assets/${this.model}.model.json`},
    //   display: {position: 'left', width: 200, height: 300},
    //   mobile: {show: true, scale: 0.5},
    // })
  }
}
</script>

<style>
</style>