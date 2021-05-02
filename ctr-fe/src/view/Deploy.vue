<template>
  <div class="split">
    <Split v-model="split">
      <div slot="left" class="split-pane">
        <ServiceTree ref="ServiceTree" @onTreeNodeSelected="onTreeSelectedOrExpand"
                     @onToggleExpend="onTreeSelectedOrExpand"></ServiceTree>
      </div>
      <div slot="right" class="split-pane">
        <router-view></router-view>
      </div>
    </Split>
  </div>
</template>

<script>
import ServiceTree from "@/components/Tree/ServiceTree";

export default {
  name: "Deploy",
  components: {ServiceTree},
  data() {
    return {
      split: 0.15
    }
  },
  methods: {
    onTreeSelectedOrExpand(node) {
      if (node.type === 4) {
        this.$router.push({name: 'RouterInfo', query: {id: node.id}})
      } else {
        console.log(node)
      }
    }
  },
}
</script>

<style scoped>
.split {
  height: 100vh;
}

.split-pane {
  padding: 10px;
}
</style>