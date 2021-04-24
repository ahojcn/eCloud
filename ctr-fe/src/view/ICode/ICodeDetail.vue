<template>
  <div v-if="!loading">
    <Row :gutter="18">
      <Col span="12">
        <Card shadow :title="icode_detail.name + ''" style="min-height: 100vh">
          <Form :label-width="80">
            <FormItem label="Web IDE">
              <a @click="onClickICodeURL('http://' + icode_detail.host_info.ip + ':' + icode_detail.port)">
                http://{{ icode_detail.host_info.ip }}:{{ icode_detail.port }}
              </a>
            </FormItem>
            <FormItem label="用户组">
              <Tooltip transfer>
                <Avatar style="color: #f56a00;background-color: #fde3cf">{{ icode_detail.user_info.username }}</Avatar>
                <div slot="content">
                  <p>邮箱：{{ icode_detail.user_info.email }}</p>
                  <p>手机：{{ icode_detail.user_info.phone }}</p>
                  <p>状态：{{ icode_detail.user_info.is_active }}</p>
                </div>
              </Tooltip>
            </FormItem>
            <FormItem label="容器id">
              <Input readonly v-model="icode_detail.container_id"></Input>
            </FormItem>
            <FormItem label="密码">
              <Input readonly type="password" password v-model="icode_detail.password"></Input>
            </FormItem>
            <FormItem label="端口映射">
              <Tag color="red">{{ icode_detail.port }}</Tag>
              ：
              <Tag color="green">{{ icode_detail.container_port }}</Tag>
            </FormItem>
            <FormItem label="创建于">
              <Time :time="new Date(icode_detail.create_time)"></Time>
            </FormItem>
            <FormItem label="其他信息">
              <span style="background-color: #fde3cf">{{ icode_detail.extra === '' ? '暂无' : icode_detail.extra }}</span>
            </FormItem>
            <FormItem>
              <Button type="error" @click="onDeleteICodeBtnClick">删除</Button>
            </FormItem>
          </Form>
        </Card>
      </Col>
      <Col span="12">
        <Card shadow :title="icode_detail.name + ' - 关联主机'" style="min-height: 100vh">
          <Form :label-width="80">
            <FormItem label="主机描述">
              <Input readonly type="textarea" v-model="icode_detail.host_info.description"></Input>
            </FormItem>
            <FormItem label="连接信息">
              <Input readonly type="text"
                     :value="icode_detail.host_info.username + '@' + icode_detail.host_info.ip + ':' + icode_detail.host_info.port"></Input>
            </FormItem>
            <FormItem label="创建于">
              <Time :time="icode_detail.host_info.create_time"></Time>
            </FormItem>
            <FormItem label="用户组">
              <Tooltip transfer v-for="(user_info, index) in icode_detail.host_info.user_list" :key="index">
                <Tag v-if="user_info.id===icode_detail.host_info.create_user.id"
                     color="red">@{{ user_info.username }}（管理员）
                </Tag>
                <Tag v-if="user_info.id!==icode_detail.host_info.create_user.id">@{{ user_info.username }}</Tag>
                <div slot="content">
                  <p>邮箱：{{ user_info.email }}</p>
                  <p>手机：{{ user_info.phone }}</p>
                  <p>状态：{{ user_info.is_active }}</p>
                </div>
              </Tooltip>
            </FormItem>
            <FormItem lebel="详细配置" style="color: slategray">
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
            </FormItem>
          </Form>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiGetICodeList, apiDeleteICode} from "@/api/icode";

export default {
  name: "ICodeDetail",
  data() {
    return {
      icode_detail: {},
      loading: true
    }
  },
  computed: {
    id: function () {
      return this.$route.query.id
    },
    host_info_extra: function () {
      return JSON.parse(this.icode_detail.host_info.extra)
    }
  },
  watch: {
    id(new_id) {
      this.getDetail(new_id)
    }
  },
  methods: {
    getDetail(id) {
      apiGetICodeList({id: id}).then(res => {
        if (res.code === 200) {
          this.icode_detail = res.data[0]
          this.loading = false
        }
      })
    },
    onClickICodeURL(url) {
      window.open(url, "iCode", "height=754, width=1277, top=0, left=2, toolbar=no, menubar=no, scrollbars=no, resizable=yes,location=no, status=no")
    },
    onDeleteICodeBtnClick() {
      this.$Modal.confirm({
        title: '请确认',
        content: `是否删除【${this.icode_detail.name}】`,
        onOk: () => {
          apiDeleteICode({id: this.id}).then(res => {
            if (res.code === 200) {
              this.$emit('refreshICodeList')
            }
          })
        },
        onCancel: () => {
          this.$Message.info('已取消')
        }
      })
    },
  },
  mounted() {
    this.getDetail(this.id)
  }
}
</script>

<style scoped>

</style>