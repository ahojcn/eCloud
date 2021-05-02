<template>
  <div>
    <Button type="primary" ghost size="small" icon="md-add" @click="show_host_user_add_modal = true">
      添加用户
    </Button>

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
import SelectUserByUsernameOrEmail from "@/components/User/SelectUserByUsernameOrEmail";
import {apiAddHostUser} from "@/api/host";

export default {
  name: "HostUserAdd",
  components: {SelectUserByUsernameOrEmail},
  props: {
    host_info: {
      required: true,
      type: Object
    }
  },
  data() {
    return {
      show_host_user_add_modal: false
    }
  },
  methods: {
    onHostUserAddUserChange(user_id) {
      this.host_user_add_user_id = user_id
    },
    onHostUserAddModalBtnClick() {
      apiAddHostUser({
        host_id: this.host_info.id,
        user_id: this.host_user_add_user_id
      }).then(res => {
        if (res.code === 200) {
          this.$emit('onAddHostUserOk')
          this.show_host_user_add_modal = false
        }
      })
    },
  }
}
</script>

<style scoped>

</style>