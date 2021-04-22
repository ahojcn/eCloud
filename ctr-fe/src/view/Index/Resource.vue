<template>
  <div>
    <AddHostModal style="text-align: right" @onAddHostOk="refreshHostList"></AddHostModal>

    <Table :loading="host_list_loading" stripe :columns="host_list_columns" :data="host_list">
      <template slot-scope="{ row }" slot="name">
        {{ row.username }}@{{ row.ip }}:{{ row.port }}
      </template>
      <template slot-scope="{row}" slot="time">
        创建于
        <Time :time="row.create_time"></Time>
        <br/>
        更新于
        <Time :time="row.update_time"></Time>
      </template>
      <template slot-scope="{row}" slot="admin">
        <Tooltip transfer>
          <Avatar style="color: #f56a00;background-color: #fde3cf">{{ row.create_user.username }}</Avatar>
          <div slot="content">
            <p>邮箱：{{ row.create_user.email }}</p>
            <p>手机：{{ row.create_user.phone }}</p>
            <p>状态：{{ row.create_user.is_active }}</p>
          </div>
        </Tooltip>
      </template>
      <template slot-scope="{row}" slot="users">
        <Tooltip transfer v-for="(user, index) in row.user_list" :key="index">
          <span>{{ user.username }},</span>
          <div slot="content">
            <p>邮箱：{{ user.email }}</p>
            <p>手机：{{ user.phone }}</p>
            <p>状态：{{ user.is_active }}</p>
            <Button type="error" size="small" @click="onHostUserDeleteBtnClick(row, user)">删除</Button>
          </div>
        </Tooltip>
        <Button type="primary" ghost size="small" icon="md-add" @click="onHostUserAddBtnClick(row)"></Button>
      </template>
      <template slot-scope="{row}" slot="extra">
        <ButtonGroup>
          <Button ghost type="primary" size="small" @click="onShowHostInfoDetailBtnClick(row)">详情</Button>
          <Button ghost type="primary" size="small" @click="onShowMonitorBtnClick(row)">监控</Button>
          <Button ghost type="error" size="small" @click="onHostDeleteBtnClick(row)">删除</Button>
        </ButtonGroup>
      </template>
    </Table>

    <Modal transfer footer-hide v-model="show_host_extra_modal">
      <div v-if="show_host_extra_modal">
        <strong>{{ host_info.username + '@' + host_info.ip + ':' + host_info.port }}</strong>
        <table>
          <tr>
            <td>主机名：</td>
            <td>{{ host_info_extra.host_info.info.hostname }}</td>
          </tr>
          <tr>
            <td>启动时间：</td>
            <td>
              <Time :time="host_info_extra.host_info.info.bootTime * 1000"></Time>
            </td>
          </tr>
          <tr>
            <td>操作系统：</td>
            <td>{{ host_info_extra.host_info.info.os }} {{ host_info_extra.host_info.info.platform }}
              {{ host_info_extra.host_info.info.platformVersion }}
            </td>
          </tr>
          <tr>
            <td>内核版本：</td>
            <td>{{ host_info_extra.host_info.info.kernelVersion }} {{ host_info_extra.host_info.info.kernelArch }}</td>
          </tr>
          <tr>
            <td>CPU：</td>
            <td>{{ host_info_extra.cpu_info.info[0].modelName }}</td>
          </tr>
          <tr>
            <td>CPU频率：</td>
            <td>{{ host_info_extra.cpu_info.info[0].mhz }}MHz</td>
          </tr>
          <tr>
            <td>CPUCache：</td>
            <td>{{ host_info_extra.cpu_info.info[0].cacheSize }}MB</td>
          </tr>
          <tr>
            <td>核心：</td>
            <td>{{ host_info_extra.cpu_info.physical }}（物理） * {{ host_info_extra.cpu_info.logical }}（逻辑）</td>
          </tr>
        </table>
      </div>
    </Modal>

    <Drawer v-model="show_monitor_drawer" width="100" scrollable mask :title="'监控信息 - ' + host_info.ip">
      <MonitorMetricsSelect @onMonitorMetricsSelected="onMonitorMetricsSelected"></MonitorMetricsSelect>
      <Chart :options="monitor_options"></Chart>
    </Drawer>

    <Modal transfer v-model="show_host_user_add_modal" title="给用户添加权限">
      <Form :label-width="80">
        <FormItem label="主机">
          {{ host_info.ip }}
        </FormItem>
        <FormItem label="用户">
          <SelectUserByUsernameOrEmail @onChange="onHostUserAddUserChange"></SelectUserByUsernameOrEmail>
        </FormItem>
        <FormItem label="注意">
          <Alert type="error">添加后此用户可查看主机相关资源 <br/> 可部署服务在此主机上</Alert>
        </FormItem>
      </Form>

      <div slot="footer">
        <Button type="primary" long @click="onHostUserAddModalBtnClick">添加</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import {apiGetHostList, apiDeleteHost, apiDeleteHostUser, apiAddHostUser} from "@/api/host";
import {apiGetMetricsData} from "@/api/monitor";
import VueShell from 'vue-shell';
import MonitorMetricsSelect from "@/components/MonitorMetricsSelect";
import Chart from "@/components/Chart";
import SelectUserByUsernameOrEmail from "@/components/SelectUserByUsernameOrEmail";
import RunCommand from '@/components/RunCommand';
import AddHostModal from "@/components/AddHostModal";

export default {
  name: "Resource",
  // eslint-disable-next-line vue/no-unused-components
  components: {AddHostModal, SelectUserByUsernameOrEmail, Chart, MonitorMetricsSelect, VueShell, RunCommand},
  props: {},
  data() {
    return {
      host_list_columns: [
        {
          type: 'expand',
          width: 50,
          render: (h, params) =>{
            return h(RunCommand, {
              props: {
                host_info: params.row
              }
            })
          }
        },
        {title: 'id', key: 'id'},
        {title: 'name', slot: 'name'},
        {title: '时间', slot: 'time'},
        {title: '管理员', slot: 'admin'},
        {title: '用户组', slot: 'users'},
        {title: '主机信息', slot: 'extra'},
      ],
      host_list: [],
      host_list_loading: true,
      show_host_extra_modal: false,
      host_info_extra: {},
      host_info: {},
      show_monitor_drawer: false,
      show_monitor_charts: false,
      monitor_options: {
        title: {
          text: '监控图'
        },
        dataZoom: [
          {
            id: 'dataZoomX',
            type: 'slider',
            xAxisIndex: [0],
            filterMode: 'filter', // 设定为 'filter' 从而 X 的窗口变化会影响 Y 的范围。
            start: 0,
            end: 100
          },
          {
            id: 'dataZoomY',
            type: 'slider',
            yAxisIndex: [0],
            filterMode: 'filter',
            start: 0,
            end: 100
          }
        ],
        toolbox: {
          show: true,
          showTitle: false, // 隐藏默认文字，否则两者位置会重叠
          feature: {
            saveAsImage: {
              show: true,
              title: 'Save As Image'
            },
            magicType: {
              type: ['line', 'bar', 'stack', 'tiled']
            },
            dataView: {
              show: true,
              title: 'Data View'
            },
            dataZoom: {
              show: true,
              title: '缩放'
            }
          },
        },
        tooltip: {
          show: true,
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
          },
        },
        legend: {
          data: []
        },
        xAxis: {
          type: 'time',
          splitLine: {
            show: false
          },
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, '100%'],
          splitLine: {
            show: false
          }
        },
        series: [],
      },
      show_host_user_add_modal: false,
      host_user_add_user_id: -1,
    }
  },
  methods: {
    onShowHostInfoDetailBtnClick(host_info) {
      this.show_host_extra_modal = true
      this.host_info_extra = JSON.parse(host_info.extra)
      this.host_info = host_info
    },
    onShowMonitorBtnClick(host_info) {
      this.show_monitor_drawer = true
      this.host_info = host_info
    },
    onMonitorMetricsSelected(metrics, cols, from_time, to_time) {
      this.monitor_options.series.length = 0
      this.monitor_options.title.text = metrics + '(' + cols + ')'
      for (let i = 0; i < cols.length; i++) {
        apiGetMetricsData({
          "host_id": this.host_info.id,
          "metrics": metrics,
          "cols": cols[i],
          "from_time": from_time,
          "to_time": to_time
        }).then(res => {
          this.monitor_options.series.push({type: "line", name: cols[i], data: res.data[0].Series[0].values})
          this.monitor_options.legend.data.push(cols[i])
        })
      }
    },
    onHostDeleteBtnClick(host_info) {
      this.$Modal.confirm({
        title: '请确认',
        content: '是否删除 <span style="color: red">ip=' + host_info.ip + '？</span>',
        onOk: () => {
          apiDeleteHost(host_info.id).then(res => {
            if (res.code === 200) {
              this.refreshHostList()
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消')
        }
      })
    },
    refreshHostList() {
      apiGetHostList().then(res => {
        if (res.code === 200) {
          this.host_list = res.data
          this.host_list_loading = false
        }
      })
    },
    onHostUserDeleteBtnClick(host_info, user_info) {
      this.$Modal.confirm({
        title: '请确认',
        content: `是否删除用户
            <span style="color: red">${user_info.username}</span>
            在主机（<span style="color: red">${host_info.ip} | ${host_info.description}</span>）的权限？`,
        onOk: () => {
          apiDeleteHostUser({host_id: host_info.id, user_id: user_info.id}).then(res => {
            if (res.code === 200) {
              this.refreshHostList()
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消')
        }
      })
    },
    onHostUserAddBtnClick(host_info) {
      this.host_info = host_info
      this.show_host_user_add_modal = true
    },
    onHostUserAddUserChange(user_id) {
      this.host_user_add_user_id = user_id
    },
    onHostUserAddModalBtnClick() {
      apiAddHostUser({
        host_id: this.host_info.id,
        user_id: this.host_user_add_user_id
      }).then(res => {
        if (res.code === 200) {
          this.refreshHostList()
          this.show_host_user_add_modal = false
        }
      })
    },
  },
  mounted() {
    this.refreshHostList()
  }
}
</script>

<style scoped>

</style>