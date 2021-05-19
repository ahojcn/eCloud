<template>
  <div>

  </div>
</template>

<script>
import {apiClusterList} from "@/api/cluster";

export default {
  name: "ClusterList",
  data() {
    return {}
  },
  methods: {},
  mounted() {
    apiClusterList({tree_id: 30}).then(res => {
      console.log(res)
    })
  }
}
</script>

<style scoped>

</style>