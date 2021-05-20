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
      <template slot="status" slot-scope="scope">
        <Tag size="large" :color="getStatusColor(scope.row.status)">
          {{ scope.row.status_msg }}
        </Tag>
      </template>
      <template slot="router" slot-scope="scope">
        <p v-if="scope.row.status===6">
          {{ scope.row.router_ip }}:{{ scope.row.router_port }}
        </p>
      </template>
      <template slot="detail" slot-scope="scope">
        <Button long size="small" icon="ios-information-circle" type="primary" @click="onAliveDetailBtnClick(scope.row)">
          存活检测信息
        </Button>
        <br>
        <Button long size="small" icon="md-play" type="success" @click="onRunPipeLineBtnClick(scope.row)">
          执行流水线
        </Button>
        <br>
        <Button long size="small" icon="md-redo" type="warning" @click="onResetPipeLineBtnClick(scope.row)">
          重置流水线
        </Button>
        <br>
        <Button long size="small" icon="md-trash" type="error" @click="onDeletePipeLineBtnClick(scope.row)">
          删除流水线
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
    <Drawer v-model="show_pipeline_run_drawer" width="100" :mask-closable="false">
      <template slot="header">
        <Button type="success" icon="md-play" style="text-align: center" @click="onRunPipeLineDrawerBtnClick">
          执行
        </Button>
        <Button type="primary" icon="md-refresh" style="text-align: center"
                @click="refreshPipeLineStatus(pipeline_run.id)">
          刷新
        </Button>
      </template>
      <Row :gutter="20">
        <Col :span="6">
          <Card style="min-height: 100vh">
            <Steps direction="vertical" :current="pipeline_status.current">
              <Step v-for="(step, index) in pipeline_status.steps" :key="index" :title="step"
                    :content="pipeline_status.content[index]"></Step>
            </Steps>
          </Card>
        </Col>
        <Col :span="18">
          <Card title="日志" :bordered="false" style="min-height: 100vh">
            <Input type="textarea" v-model="pipeline_status.logs" :border="false" autosize readonly
                   placeholder="日志"></Input>
          </Card>
        </Col>
      </Row>
    </Drawer>
  </div>
</template>

<script>
import {
  apiDeletePipeLine,
  apiGetPipeLineList,
  apiGetPipeLineStatus,
  apiResetPipeLine,
  apiRunPipeLine
} from "@/api/pipeline";

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
        {title: '状态', slot: 'status', align: 'center'},
        {title: '接入层入口', slot: 'router', align: 'center'},
        {title: '集群id', key: 'cluster_id'},
        {title: '操作', slot: 'detail', width: 200},
      ],
      alive_detail: {},
      show_alive_detail_modal: false,

      pipeline_run: {},
      pipeline_status: {},
      show_pipeline_run_drawer: false,
    }
  },
  watch: {
    show_pipeline_run_drawer: function () {
      this.getData()
    },
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
      this.pipeline_run = row
      apiGetPipeLineStatus({id: row.id}).then(res => {
        this.pipeline_status = res.data
        this.show_pipeline_run_drawer = true
      })
    },
    onRunPipeLineDrawerBtnClick() {
      apiRunPipeLine({id: this.pipeline_run.id}).then(res => {
        if (res.code === 200) {
          this.refreshPipeLineStatus(this.pipeline_run.id)
        }
      })
    },
    refreshPipeLineStatus(id) {
      apiGetPipeLineStatus({id: id}).then(res => {
        this.pipeline_status = res.data
        console.log(res.data)
      })
    },
    onResetPipeLineBtnClick(row) {
      this.$Modal.confirm({
        title: '此操作不可恢复，请谨慎操作',
        content: '<p>1、集群配置中当前容器数归 0</p>' +
            '<p>2、删除集群相关联的容器</p>' +
            '<p>3、更新流水线状态=Init</p>' +
            '<p>4、删除接入层配置</p>' +
            '<p>5、重启接入层</p>',
        onOk: () => {
          apiResetPipeLine({id: row.id}).then(res => {
            if (res.code === 200) {
              this.$Message.success('重置流水线完成')
              this.getData()
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消重置');
        }
      });
    },
    onDeletePipeLineBtnClick(row) {
      this.$Modal.confirm({
        title: '此操作不可恢复，请谨慎操作',
        content: '<p>1、集群配置中当前容器数归 0</p>' +
            '<p>2、删除集群相关联的容器</p>' +
            '<p>3、更新流水线状态=Init</p>' +
            '<p>4、删除接入层配置</p>' +
            '<p>5、重启接入层</p>' +
            '<p>6、删除流水线</p>',
        onOk: () => {
          apiDeletePipeLine({id: row.id}).then(res => {
            if (res.code === 200) {
              this.getData()
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消删除');
        }
      });
    },
    getStatusColor(status) {
      switch (status) {
        case 0:
          return 'error'
        case 1:
          return 'primary'
        case 2:
          return 'cyan'
        case 3:
          return 'blue'
        case 4:
          return 'purple'
        case 5:
          return 'green'
        case 6:
          return 'success'
      }
    },
  }
}
</script>

<style scoped>

</style>