<template>
  <div>
    <div v-if="show_router_info">
      <Card shadow>
        <template slot="title">
          <h3>接入层信息：UN={{tree_node_info.user_tree_info.un}}</h3>
        </template>
        <template slot="extra">
          <Button type="primary" icon="md-redo" @click="onRouterRedoBtnClick">重新部署接入层</Button>
        </template>
        <Form :label-width="200">
          <FormItem label="主机信息">
            {{ router_info.host_info.username }}@{{ router_info.host_info.ip }}:{{ router_info.host_info.port }}
            <Button :to="{name: 'HostDetail', query: {id: router_info.host_info.id}}" type="primary" size="small"
                    icon="md-arrow-forward"></Button>
          </FormItem>
          <FormItem label="服务树信息">
            {{ router_info.ns_info.un }}
            <Button :to="{name: 'TreeNodeDetail', query: {id: router_info.ns_info.id}}" type="primary" size="small"
                    icon="md-arrow-forward"></Button>
            <br>
            <Alert type="info">
              {{ router_info.ns_info.description }}
            </Alert>
          </FormItem>
          <FormItem label="接入层状态">
            <Input autosize readonly type="textarea" v-model="router_status.nginx_status"></Input>
          </FormItem>
          <FormItem label="logstash状态">
            <Input autosize readonly type="textarea" v-model="router_status.logstash_status"></Input>
          </FormItem>
          <FormItem label="logstash配置">
            <Input autosize readonly type="textarea" v-model="router_status.logstash_config"></Input>
          </FormItem>
          <FormItem label="接入层信息">
            创建于
            <Time :time="router_info.router.create_time"></Time>
            ，更新于
            <Time :time="router_info.router.update_time"></Time>
            <Input :rows="30" readonly type="textarea" v-model="router_info.router.log"></Input>
          </FormItem>
        </Form>
      </Card>
    </div>
    <div v-else>
      <h1>
        暂未部署接入层
        <Button type="primary" @click="show_router_deploy_modal = true">部署</Button>
      </h1>

      <Input readonly type="textarea" :rows="50" v-model="deploy_log" placeholder="部署日志"></Input>
      <Modal v-model="show_router_deploy_modal" title="接入层部署">
        <Form :label-width="80">
          <FormItem v-if="show_tree_node_info" label="服务树NS">
            {{ tree_node_info.user_tree_info.un }}
          </FormItem>
          <FormItem label="主机">
            <SelectHost @onHostSelected="onHostSelected"></SelectHost>
          </FormItem>
        </Form>

        <template slot="footer">
          <Button long type="primary" @click="handleRouterDeployBtnClick">部署</Button>
        </template>
      </Modal>
    </div>
  </div>
</template>

<script>
import {apiGetTreeInfo} from "@/api/tree";
import {apiGetRouterInfo, apiGetRouterStatus, apiRouterDeploy, apiRouterRedo} from "@/api/m_router";

import SelectHost from "@/components/Host/SelectHost";

export default {
  name: "Router",
  components: {SelectHost},
  computed: {
    id: function () {
      return this.$route.query.id
    }
  },
  watch: {
    id(new_id) {
      this.refreshTreeNodeInfo(new_id)
      this.refreshRouterInfo({ns_id: new_id})
    },
  },
  mounted() {
    this.refreshTreeNodeInfo(this.id)
    this.refreshRouterInfo({ns_id: this.id})
  },
  data() {
    return {
      tree_node_info: {},
      router_info: {},
      show_router_info: false,
      show_tree_node_info: false,
      show_router_deploy_modal: false,
      selected_host_id: -1,
      deploy_log: '',

      router_status: {},
    }
  },
  methods: {
    refreshTreeNodeInfo(id) {
      apiGetTreeInfo({"id": id}).then(res => {
        if (res.code === 200) {
          this.tree_node_info = res.data
          this.show_tree_node_info = true
        } else {
          this.show_tree_node_info = false
        }
      })
    },
    refreshRouterInfo(obj) {
      apiGetRouterInfo(obj).then(res => {
        if (res.code === 200) {
          if (res.data.length !== 0) {
            this.show_router_info = true
            this.router_info = res.data[0]
            this.getRouterStatus(this.router_info.router.id)
          } else {
            this.show_router_info = false
          }
        }
      })
    },
    onHostSelected(host_id) {
      this.selected_host_id = host_id
    },
    handleRouterDeployBtnClick() {
      if (this.selected_host_id < 0) {
        this.$Modal.error('请选择一台主机用于部署接入层')
      } else {
        this.$Spin.show()
        apiRouterDeploy({ns_id: this.tree_node_info.user_tree_info.id, host_id: this.selected_host_id})
            .then(res => {
              if (res.code === 200) {
                this.deploy_log = res.data[0]
                this.show_router_deploy_modal = false
              }
              this.$Spin.hide()
            })
      }
    },
    getRouterStatus(id) {
      apiGetRouterStatus({id: id}).then(res => {
        this.router_status = res.data
      })
    },
    onRouterRedoBtnClick() {
      this.$Modal.confirm({
        title: '此操作不可恢复，请谨慎操作',
        content: '<p>1、删除所有router的配置文件 0</p>' +
            '<p>2、删除所有相关联的pipeline</p>',
        onOk: () => {
          apiRouterRedo({router_id: this.router_info.router.id}).then(res => {
            if (res.code === 200) {
              this.refreshRouterInfo()
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消重置');
        }
      });
    },
  },
}
</script>

<style scoped>

</style>