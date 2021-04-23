<template>
  <div>
    <Row :gutter="16">
      <Col span="4">
        <ServiceTree ref="ServiceTree" @onTreeNodeSelected="onTreeSelectedOrExpand"
                     @onToggleExpend="onTreeSelectedOrExpand"></ServiceTree>
      </Col>
      <Col span="20">
        <router-view @onAddUserTreeSuccess="refreshTree"
                     @onTreeNodeAddSuccessful="refreshTree"></router-view>
      </Col>
    </Row>
  </div>
</template>

<script>
import ServiceTree from "@/components/ServiceTree";

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
}
</script>

<style scoped>

</style>