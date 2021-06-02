<template>
  <div>
    <Button type="primary" v-if="show_service_config===false" :disabled="show_service_config"
            @click="show_service_config_add_model = true">新增预案
    </Button>

    <Card v-else>
      <template slot="title">
        <h3>
          预案配置
          <Tag color="green">ip={{ service_info.router_ip }}</Tag>
          <Tag color="red">port={{ service_info.router_port }}</Tag>
        </h3>
      </template>
      <template slot="extra">
        <Button type="error" :disabled="show_service_config===false" @click="onDeleteServiceConfigBtnClick">删除预案
        </Button>
      </template>
      <Row>
        <Col :span="12">
          <Chart style="height: 300px" :options="default_chart_option" id="tree"></Chart>
        </Col>
        <Col :span="12">
          <Input autosize readonly type="textarea" v-model="service_info.config_content"></Input>
        </Col>
      </Row>
    </Card>

    <Modal title="创建分流预案" v-model="show_service_config_add_model">
      <Form :label-width="80">
        <FormItem v-for="(item, index) in service_config_add_form.flow_map" :key="index"
                  :label="'cluster' + index">
          <Row>
            <Col span="18">
              <ClusterSelect :tree_id="tree_id" :index="index"
                             @onClusterSelected="handleClusterSelected"></ClusterSelect>
              <InputNumber :formatter="value => `${value}%`"
                           :step="5"
                           :parser="value => value.replace('%', '')"
                           controls-outside :max="100" :min="0" v-model="item.flow_percent"></InputNumber>
            </Col>
            <Col span="4">
              <Button @click="handleDeleteClusterItem(index)">Delete</Button>
            </Col>
          </Row>
        </FormItem>
        <FormItem>
          <Row>
            <Col span="12">
              <Button type="dashed" long @click="handleAddItem" icon="md-add">Add item</Button>
            </Col>
          </Row>
        </FormItem>
        <FormItem>
          <Button type="primary" long @click="onAddServiceConfigBtnClick">提交</Button>
        </FormItem>
      </Form>
      <div slot="footer">
        创建后预案将自动生效
      </div>
    </Modal>

  </div>
</template>

<script>
import {apiCreateService, apiDeleteService, apiGetService} from "@/api/service";
import Chart from "@/components/Chart";
import ClusterSelect from "@/components/Cluster/Select";

export default {
  name: "ServiceConfig",
  components: {ClusterSelect, Chart},
  props: {
    tree_id: {
      required: true,
      type: Number,
    }
  },
  data() {
    return {
      service_info: {},
      show_service_config: false,

      show_service_config_add_model: false,
      service_config_add_form: {
        flow_map: [
          {cluster_id: '', flow_percent: 0}
        ],
      },

      default_chart_option: {
        tooltip: {
          trigger: 'item',
          triggerOn: 'mousemove',
          formatter: 'UN：<strong>{b}</strong><br/>流量占比：<strong>{c}%</strong>',
        },
        series: [
          {
            type: 'tree',
            data: [],
            orient: 'vertical',
            label: {
              position: 'bottom',
              verticalAlign: 'middle',
              align: 'center',
              fontSize: 15
            },
            leaves: {
              label: {
                position: 'bottom',
                verticalAlign: 'middle',
                align: 'center'
              }
            },
            emphasis: {
              focus: 'descendant'
            },
            expandAndCollapse: true,
            animationDuration: 550,
            animationDurationUpdate: 750
          }
        ]
      },
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    getData() {
      apiGetService({tree_id: this.tree_id}).then(res => {
        if (res.code === 200) {
          this.service_info = res.data
          this.show_service_config = true
          this.default_chart_option.series[0].data.push(this.service_info.chart_opt)
        } else {
          this.show_service_config = false
        }
      })
    },
    onDeleteServiceConfigBtnClick() {
      apiDeleteService({tree_id: this.tree_id}).then(res => {
        if (res.code === 200) {
          this.getData()
        }
      })
    },
    handleDeleteClusterItem(index) {
      this.service_config_add_form.flow_map = this.service_config_add_form.flow_map.slice(index, index + 1)
    },
    handleClusterSelected(v, i) {
      this.service_config_add_form.flow_map[i].cluster_id = String(v)
    },
    handleAddItem() {
      this.service_config_add_form.flow_map.push({cluster_id: '', flow_percent: 0})
    },
    onAddServiceConfigBtnClick() {
      let percent_count = 0
      let d = {}
      for (let i = 0; i < this.service_config_add_form.flow_map.length; i++) {
        let fm = this.service_config_add_form.flow_map[i]
        if (fm.cluster_id === '') {
          this.$Message.error(`参数错误，请选择集群，cluster${i}`)
          return
        }
        percent_count += fm.flow_percent
        d[fm.cluster_id] = fm.flow_percent
      }
      if (percent_count !== 100) {
        this.$Message.error(`参数错误，分流预案综合必须等于100`)
        return
      }

      apiCreateService({tree_id: this.tree_id, flow_map: JSON.stringify(d)}).then(res => {
        if (res.code === 200) {
          this.getData()
          this.$emit('onCreateServiceConfigOk')
          this.show_service_config_add_model = false
        }
      })
    },
  },
}
</script>

<style scoped>

</style>