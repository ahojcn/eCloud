<template>
  <div>
    <Form>
      <FormItem label="预期容器数量">
        <InputNumber v-model="form.cluster_num"></InputNumber>
      </FormItem>
      <FormItem label="容器端口">
        <InputNumber v-model="form.container_port"></InputNumber>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="onClusterCreateBtnClick">保存集群配置</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
import {apiClusterCreate} from "@/api/cluster";

export default {
  name: "ClusterCreate",
  data() {
    return {
      form: {
        tree_id: 0,
        cluster_num: 0,
        container_port: 8080,
      }
    }
  },
  computed: {
    tree_id: function () {
      return parseInt(this.$route.query.id)
    }
  },
  mounted() {
    this.form.tree_id = this.tree_id
  },
  methods: {
    onClusterCreateBtnClick() {
      apiClusterCreate(this.form).then(res => {
        if (res.code === 200) {
          this.$emit('onClusterCreated', res.data)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>