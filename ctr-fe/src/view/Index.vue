<template>
  <div>
    <Menu mode="horizontal" theme="light" active-name="1" @on-select="handleMenuSelect">
      <MenuItem name="1">
        <Icon type="ios-paper"/>
        内容管理
      </MenuItem>
      <MenuItem name="2">
        <Icon type="ios-people"/>
        用户管理
      </MenuItem>
      <Submenu name="3">
        <template slot="title">
          <Icon type="ios-stats"/>
          统计分析
        </template>
        <MenuGroup title="使用">
          <MenuItem name="3-1">新增和启动</MenuItem>
          <MenuItem name="3-2">活跃分析</MenuItem>
          <MenuItem name="3-3">时段分析</MenuItem>
        </MenuGroup>
        <MenuGroup title="留存">
          <MenuItem name="3-4">用户留存</MenuItem>
          <MenuItem name="3-5">流失用户</MenuItem>
        </MenuGroup>
      </Submenu>
      <MenuItem name="4">
        <Icon type="ios-construct"/>
        综合设置
      </MenuItem>

      <Submenu name="avatar">
        <template slot="title">
          <Avatar size="large" style="color: #f56a00;background-color: #fde3cf">{{ user_info.username }}</Avatar>
        </template>
        <MenuItem name="logout">
          <Icon type="md-exit" />
          退出
        </MenuItem>
      </Submenu>
    </Menu>

    <ServiceTree @onTreeNodeSelected="onTreeNodeSelected"></ServiceTree>
  </div>
</template>

<script>
import {is_login, logout} from "@/api/session";
import ServiceTree from "@/components/ServiceTree";

export default {
  name: "Index",
  components: {ServiceTree},
  computed: {
    user_info() {
      return this.$store.state.user_info
    }
  },
  mounted() {
    is_login().then(res => {
      if (res.code !== 200) {
        this.$router.push("/login")
      }
      this.$store.commit('set_user_info', res.data)
    })
  },
  methods: {
    handleMenuSelect(name) {
      if (name === 'logout') {
        this.handleClickLogout()
      }
    },
    handleClickLogout() {
      logout().then(res => {
        if (res.code === 200) {
          this.$router.go(0)
        }
      })
    },
    onTreeNodeSelected(node) {
      console.log(node)
    },
  }
}
</script>

<style scoped>

</style>