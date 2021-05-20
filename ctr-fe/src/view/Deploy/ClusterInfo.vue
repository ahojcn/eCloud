<template>
  <div>
    <Card title="集群配置" style="min-height: 100vh">
      <template slot="extra">
        <Button type="error" icon="md-delete" @click="onClusterDeleteBtnClick">删除集群配置</Button>
      </template>
      <div v-if="show_create_cluster_form === true">
        <Alert type="error">
          此集群还没有配置
        </Alert>
        <ClusterCreate @onClusterCreated="onClusterCreated" v-if="show_create_cluster_form === true"></ClusterCreate>
      </div>

      <div v-else>
        <Form inline>
          <FormItem label="容器端口">
            <InputNumber v-model="cluster_info.container_port" readonly></InputNumber>
          </FormItem>
          <FormItem label="集群容器数量">
            <InputNumber v-model="cluster_info.cluster_num" readonly></InputNumber>
          </FormItem>
          <FormItem label="当前容器数量">
            <InputNumber v-model="cluster_info.current_cluster_num" readonly></InputNumber>
          </FormItem>
        </Form>

        <Table :data="cluster_info.containers" :columns="cluster_info_table_columns">
          <template slot="host_info" slot-scope="scope">
            {{ scope.row.host_info.ip }}
            <Tag v-if="scope.row.host_info.router" color="blue">接入层</Tag>
          </template>
        </Table>
      </div>
    </Card>
  </div>
</template>

<script>
import {apiClusterDelete, apiClusterRetrieve} from "@/api/cluster";
import ClusterCreate from "@/components/Cluster/Create";

export default {
  name: "DeployClusterInfo",
  components: {ClusterCreate},
  computed: {
    tree_id: function () {
      return parseInt(this.$route.query.id)
    }
  },
  watch: {
    tree_id: function () {
      this.getClusterData()
    }
  },
  data() {
    return {
      show_create_cluster_form: false,
      cluster_info: {},

      cluster_info_table_columns: [
        {title: 'id', key: 'container_id'},
        {title: 'ip', key: 'container_ip'},
        {title: 'port', key: 'container_port'},
        {title: '主机信息', slot: 'host_info'},
        {title: '主机端口', key: 'host_port'},
        {title: '创建时间', key: 'create_time'},
      ],
    }
  },
  mounted() {
    this.getClusterData()
  },
  methods: {
    getClusterData() {
      apiClusterRetrieve({tree_id: this.tree_id}).then(res => {
        if (res.code === 404) {
          this.show_create_cluster_form = true
        } else {
          this.show_create_cluster_form = false
          this.cluster_info = res.data
        }
      })
    },
    onClusterCreated(data) {
      this.$Message.info(`当前集群容器数量：${data.current_cluster_num}`)
      this.getClusterData()
    },
    onClusterDeleteBtnClick() {
      apiClusterDelete({tree_id: this.tree_id}).then(res=>{
        if (res.code === 200) {
          this.getClusterData()
        }
      })
    },
  },
}
</script>

<style scoped>

</style>