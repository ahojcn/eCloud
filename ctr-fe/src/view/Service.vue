<template>
  <div>
    <Split v-model="split" class="split">
      <div slot="left" class="split-pane">
        <ServiceTree ref="ServiceTree" @onTreeNodeSelected="onTreeSelectedOrExpand"
                     @onToggleExpend="onTreeSelectedOrExpand"></ServiceTree>
      </div>
      <div slot="right" class="split-pane">
        <router-view @onAddUserTreeSuccess="refreshTree"
                     @onTreeNodeAddSuccessful="refreshTree"></router-view>
      </div>
    </Split>
  </div>
</template>

<script>
import ServiceTree from "@/components/Tree/ServiceTree";

export default {
  name: "Service",
  components: {ServiceTree},
  methods: {
    onTreeSelectedOrExpand(node) {
      this.$router.push({name: 'TreeNodeDetail', query: {id: node.id}})
    },
    onTreeNodeAddSuccessful() {
      this.$refs.ServiceTree.refreshTree()
    },
    refreshTree() {
      this.$refs.ServiceTree.refreshTree()
    }
  },
  data() {
    return {
      split: 0.15
    }
  }
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