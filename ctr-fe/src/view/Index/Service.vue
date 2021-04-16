<template>
  <div>
    <Row :gutter="16">
      <Col span="4">
        <ServiceTree @onTreeNodeSelected="onTreeNodeSelected"></ServiceTree>
      </Col>
      <Col span="20">
        <h1>展示服务树节点信息</h1>
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
    onTreeNodeSelected(node) {
      console.log(node)
    },
  }
}
</script>

<style scoped>

</style>