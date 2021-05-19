<template>
  <div>
    <Form>
      <FormItem label="服务树id">
        <Input v-model="form.tree_id" disabled></Input>
      </FormItem>
      <FormItem label="集群">
        <ClusterSelect @onClusterSelected="onClusterSelected" :tree_id="this.tree_id"></ClusterSelect>
      </FormItem>
      <FormItem label="镜像">
        <Input type="textarea" autosize v-model="form.container_image"></Input>
      </FormItem>
      <FormItem label="存活测试方法">
        <Input v-model="form.alive_method"></Input>
      </FormItem>
      <FormItem label="存活测试uri">
        <Input v-model="form.alive_uri"></Input>
      </FormItem>
      <FormItem label="存活测试query">
        <Input type="textarea" autosize v-model="form.alive_req_query"></Input>
      </FormItem>
      <FormItem label="存活测试body">
        <Input type="textarea" autosize v-model="form.alive_req_body"></Input>
      </FormItem>
      <FormItem label="存活测试请求头">
        <Input type="textarea" autosize v-model="form.alive_req_header"></Input>
      </FormItem>

      <FormItem label="存活测试响应状态码">
        <InputNumber v-model="form.alive_resp_status"></InputNumber>
      </FormItem>
      <FormItem label="存活测试响应体">
        <Input type="textarea" autosize v-model="form.alive_resp_body"></Input>
      </FormItem>
    </Form>

    <div slot="footer">
      <Button type="primary" long @click="onCreatePipeLineBtnClick">创建</Button>
    </div>
  </div>
</template>

<script>
import {apiCreatePipeLine} from "@/api/pipeline";
import ClusterSelect from "@/components/Cluster/Select";

export default {
  name: "PipeLineCreate",
  components: {ClusterSelect},
  props: {
    tree_id: {
      required: true,
      type: Number,
    }
  },
  mounted() {
    this.form.tree_id = this.tree_id
  },
  data() {
    return {
      form: {
        tree_id: 0,
        cluster_id: 0,
        container_image: 'luksa/kubia:latest',
        alive_method: 'GET',
        alive_uri: '/keepalived',
        alive_req_query: 'a=1&b=2',
        alive_req_body: '{"username":"ahojcn", "password":"text"}',
        alive_req_header: '{"content-type": "application/json"}',
        alive_resp_status: 200,
        alive_resp_body: '{"status": 200, "data": "ok", "msg": "成功"}',
      }
    }
  },
  methods: {
    onCreatePipeLineBtnClick() {
      apiCreatePipeLine(this.form).then(res => {
        if (res.code === 200) {
          this.$emit('onPipeLineCreated', res.data)
        }
      })
    },
    onClusterSelected(id) {
      this.$Message.success('关联集群id已修改')
      this.form.cluster_id = id
    },
  }
}
</script>

<style scoped>

</style>