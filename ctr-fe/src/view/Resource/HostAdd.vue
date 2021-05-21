<template>
  <div>
    <Card shadow title="添加主机" style="min-height: 100vh">
      <template slot="extra">
        <Button type="primary" @click="onAddHostBtnClick">添加</Button>
      </template>
      <Form :label-width="80">
        <FormItem label="描述">
          <Input v-model="form.description" type="textarea" show-word-limit maxlength="1024"
                 placeholder="主机基本信息描述"></Input>
        </FormItem>
        <FormItem label="ip">
          <Input v-model="form.ip" placeholder="ipv4地址"></Input>
        </FormItem>
        <FormItem label="用户名">
          <Input v-model="form.username" placeholder="登录用户名"></Input>
        </FormItem>
        <FormItem label="密码">
          <Input v-model="form.password" type="password" placeholder="将会被加密存储"></Input>
        </FormItem>
        <FormItem label="端口">
          <InputNumber v-model="form.port" :min="1" :max="65534" placeholder="ssh端口号"></InputNumber>
        </FormItem>
      </Form>
    </Card>
  </div>
</template>

<script>
import {apiAddHost} from "@/api/host";

export default {
  name: "HostAdd",
  data() {
    return {
      form: {
        description: '',
        ip: '',
        username: '',
        password: '',
        port: 22
      },
    }
  },
  methods: {
    onAddHostBtnClick() {
      this.$Spin.show()
      apiAddHost(this.form).then(res => {
        if (res.code === 200) {
          this.show_modal = false
          this.$emit('onAddHostOk')
        }
        this.$Spin.hide()
      })
    }
  }
}
</script>

<style scoped>

</style>