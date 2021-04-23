<template>
  <div>
    <Row :gutter="18">
      <Col span="4">
        <Steps :current="current_step" direction="vertical">
          <Step title="添加主机" content="添加一台主机用户部署服务和开发机"></Step>
          <Step title="填写表单" content="开发机名称"></Step>
          <Step title="选择合适的主机" content="构建对应的镜像"></Step>
          <Step title="启动容器" content="开启端口映射"></Step>
          <Step title="运行" content="开发机列表中查看刚创建的开发机"></Step>
        </Steps>
      </Col>
      <Col span="20" :style="{textAlign: 'center'}">
        <Card title="创建开发机" shadow :style="{minHeight: '100vh'}">
          <Input v-model="name" placeholder="开发机名称" style="width: 300px"></Input>

          <Button type="primary" style="margin-left: 10px" @click="onCreateICodeBtnClick">创建</Button>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiCreateICode} from '@/api/icode';
import {apiGetHostList} from '@/api/host';

export default {
  name: "ICodeCreate",
  data() {
    return {
      name: '',
      current_step: 0
    }
  },
  mounted() {
    apiGetHostList().then(res => {
      // 判断用户是否拥有主机
      if (res.code === 200 && res.data.length === 0) {
        this.$Modal.warning({
          title: '你现在还没有任何一个主机的权限',
          content: '请在「资源」中添加一台主机，或者让其他用户为您添加主机权限。'
        })
      } else {
        this.current_step += 1
      }
    })
  },
  methods: {
    onCreateICodeBtnClick() {
      if (this.name === '') {
        this.$Modal.error('开发机名称不能为空')
        return
      }
      this.$Spin.show()
      apiCreateICode({"name": this.name}).then(res => {
        if (res.code === 200) {
          this.current_step = 3
          this.$emit('refreshICodeList')
        }
        this.$Spin.hide()
      })
    }
  }
}
</script>

<style scoped>

</style>