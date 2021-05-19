<template>
  <div>
    <Table :data="d" :columns="columns">
      <template slot-scope="scope" slot="tree">
        {{ scope.row.tree_info.un }}
      </template>
      <template slot-scope="scope" slot="cluster_info">
        {{ scope.row.cluster_info.current_cluster_num }} / {{ scope.row.cluster_info.cluster_num }}
        <br>
        现有 / 目标数
      </template>
      <template slot="detail" slot-scope="scope">
        <Button size="small" icon="ios-information-circle" type="primary" @click="onAliveDetailBtnClick(scope.row)">
          存活检测信息
        </Button>
        <Button size="small" icon="md-play" type="success" @click="onRunPipeLineBtnClick(scope.row)">
          执行流水线
        </Button>
      </template>
    </Table>

    <Modal title="存货检测配置" v-model="show_alive_detail_modal">
      <Form>
        <FormItem label="URI">
          <Input readonly v-model="alive_detail.alive_uri"></Input>
        </FormItem>
        <FormItem label="Method">
          <Input readonly v-model="alive_detail.alive_method"></Input>
        </FormItem>
        <FormItem label="请求Body">
          <Input readonly type="textarea" autosize v-model="alive_detail.alive_req_body"></Input>
        </FormItem>
        <FormItem label="请求Header">
          <Input readonly type="textarea" autosize v-model="alive_detail.alive_req_header"></Input>
        </FormItem>
        <FormItem label="请求Query">
          <Input readonly type="textarea" autosize v-model="alive_detail.alive_req_query"></Input>
        </FormItem>
        <FormItem label="响应Body">
          <Input readonly type="textarea" autosize v-model="alive_detail.alive_resp_body"></Input>
        </FormItem>
        <FormItem label="响应StatusCode">
          <InputNumber readonly v-model="alive_detail.alive_resp_status"></InputNumber>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import {apiGetPipeLineList} from "@/api/pipeline";

export default {
  name: "PipeLineList",
  props: {
    tree_id: {
      required: true,
      type: Number,
    }
  },
  data() {
    return {
      d: [],
      columns: [
        {title: '服务树un', slot: 'tree', align: 'center'},
        {title: '容器', slot: 'cluster_info'},
        {title: '状态', key: 'status_msg'},
        {title: '集群id', key: 'cluster_id'},
        {title: '操作', slot: 'detail'},
      ],
      alive_detail: {},
      show_alive_detail_modal: false,

      pipeline_run: {},
      show_pipeline_run_drawer: false,
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    getData() {
      apiGetPipeLineList({tree_id: this.tree_id}).then(res => {
        this.d = res.data
      })
    },
    onAliveDetailBtnClick(row) {
      this.alive_detail = row
      this.show_alive_detail_modal = true
    },
    onRunPipeLineBtnClick(row) {
      console.log(row)
      this.pipeline_run = row
      this.show_pipeline_run_drawer = true
    }
  }
}
</script>

<style scoped>

</style>