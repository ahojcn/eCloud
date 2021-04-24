<template>
  <div v-if="!loading">
    <Card style="min-height: 100vh" shadow :title="host_info.ip" icon="md-cloud-done">
      <Row :gutter="18">
        <Col span="6">
          <Card title="配置清单">
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
                <td>{{ host_info_extra.host_info.info.kernelVersion }} {{
                    host_info_extra.host_info.info.kernelArch
                  }}
                </td>
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
          </Card>
        </Col>
        <Col span="18">
          <Form :label-width="80">
            <FormItem label="描述信息">
              <Input readonly type="textarea" v-model="host_info.description"></Input>
            </FormItem>
            <FormItem label="连接信息">
              <Input style="width: 200px" readonly type="text" v-model="host_info.username"></Input>
              <Input style="width: 200px" readonly type="text" v-model="host_info.ip"></Input>
              <Input style="width: 200px" readonly type="text" v-model="host_info.port"></Input>
            </FormItem>
            <FormItem label="用户组">
              <HostUserAdd @onAddHostUserOk="onAddHostUserOk" :host_info="host_info"></HostUserAdd>
              <Tooltip transfer v-for="(user_info, index) in host_info.user_list" :key="index">
                <Tag :color="user_info.id === host_info.create_user.id ? 'red' : 'blue'">
                  @{{ user_info.username }}
                  {{ user_info.id === host_info.create_user.id ? '（管理员）' : '' }}
                </Tag>
                <div slot="content">
                  <p>邮箱：{{ user_info.email }}</p>
                  <p>手机：{{ user_info.phone }}</p>
                  <p>状态：{{ user_info.is_active }}</p>
                  <Button type="error" size="small" @click="onHostUserDeleteBtnClick(user_info)">删除</Button>
                </div>
              </Tooltip>
            </FormItem>
          </Form>
        </Col>
      </Row>
    </Card>
  </div>
</template>

<script>
import {apiDeleteHostUser, apiGetHostList} from '@/api/host';
import HostUserAdd from "@/view/Resource/HostUserAdd";

export default {
  name: "HostDetail",
  components: {HostUserAdd},
  data() {
    return {
      host_info: {},
      loading: true,
    }
  },
  computed: {
    id: function () {
      return this.$route.query.id
    },
    host_info_extra: function () {
      return JSON.parse(this.host_info.extra)
    },
  },
  watch: {
    id(new_id) {
      this.getHostInfo(new_id)
    },
  },
  methods: {
    getHostInfo(id) {
      apiGetHostList({"id": id}).then(res => {
        if (res.code === 200) {
          this.host_info = res.data[0]
          this.loading = false
        }
      })
    },
    onHostUserDeleteBtnClick(user_info) {
      this.$Modal.confirm({
        title: '请确认',
        content: `是否删除用户
            <span style="color: red">${user_info.username}</span>
            在主机（<span style="color: red">${this.host_info.ip} | ${this.host_info.description}</span>）的权限？`,
        onOk: () => {
          apiDeleteHostUser({host_id: this.host_info.id, user_id: user_info.id}).then(res => {
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
    onAddHostUserOk() {
      this.$emit('onAddHostUserOk')
    },
  },
  mounted() {
    this.getHostInfo(this.id)
  },
}
</script>

<style scoped>

</style>