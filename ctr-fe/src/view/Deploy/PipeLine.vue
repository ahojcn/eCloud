<template>
  <div>
    <ServiceConfig :tree_id="parseInt($route.query.id)"></ServiceConfig>

    <Card style="min-height: 100vh">
      <template slot="title">
        <h3>
          流水线信息
          <Button type="text" icon="md-refresh" @click="$refs.pipeline_list.getData()"></Button>
        </h3>
      </template>
      <template slot="extra">
        <Button type="primary" @click="show_create_pipeline_modal = true">新建流水线</Button>
      </template>
      <div>
        <PipeLineList ref="pipeline_list" :tree_id="parseInt($route.query.id)"></PipeLineList>
      </div>
    </Card>

    <Modal v-model="show_create_pipeline_modal" title="创建流水线">
      <PipeLineCreate @onPipeLineCreated="onPipeLineCreated" :tree_id="parseInt($route.query.id)"></PipeLineCreate>
      <template slot="footer">创建流水线后会自动刷新</template>
    </Modal>
  </div>
</template>

<script>
import PipeLineList from "@/components/PipeLine/List";
import PipeLineCreate from "@/components/PipeLine/Create";
import ServiceConfig from "@/components/Service/ServiceConfig";

export default {
  name: "PipeLine",
  components: {ServiceConfig, PipeLineCreate, PipeLineList},
  data() {
    return {
      show_create_pipeline_modal: false,
    }
  },
  methods: {
    onPipeLineCreated() {
      this.show_create_pipeline_modal = false
      this.$refs.pipeline_list.getData()
    },
  }
}
</script>

<style scoped>

</style>