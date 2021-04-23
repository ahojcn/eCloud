<template>
  <Layout>
    <Header :style="{position: 'fixed', width: '100%', zIndex: 99}">
      <Menu :active-name="active_name" mode="horizontal" theme="dark" @on-select="handleMenuSelect">
        <div>
          <Submenu name="service">
            <template slot="title">
              <Icon type="md-infinite"/>
              服务
            </template>
            <MenuGroup title="服务树">
              <MenuItem name="service" to="/service">服务树信息</MenuItem>
              <MenuItem name="service_create" to="/service/create">创建树节点</MenuItem>
            </MenuGroup>
          </Submenu>
          <MenuItem name="monitor">
            <Icon type="ios-stats"/>
            监控
          </MenuItem>
          <MenuItem name="deploy">
            <Icon type="md-cloud-upload"/>
            部署
          </MenuItem>
          <MenuItem name="resource" to="/resource">
            <Icon type="md-code-working"/>
            资源
          </MenuItem>
          <Submenu name="dev">
            <template slot="title">
              <Icon type="ios-bug"/>
              开发
            </template>
            <MenuGroup title="开发机">
              <MenuItem name="icode_list" to="/icode">我的开发机</MenuItem>
              <MenuItem name="icdoe_create" to="/icode/create">申请开发机</MenuItem>
            </MenuGroup>
          </Submenu>
        </div>

        <div style="margin-right: 12px;float: right">
          <Submenu name="avatar">
            <template slot="title">
              <Avatar size="large" style="color: #f56a00;background-color: #fde3cf">{{ user_info.username }}</Avatar>
            </template>
            <MenuItem name="logout">
              <Icon type="md-exit"/>
              退出
            </MenuItem>
          </Submenu>
        </div>
      </Menu>
    </Header>
    <Content :style="{margin: '88px 20px 0', height: '100vh', background: '#fff'}">
      <router-view></router-view>
    </Content>
  </Layout>
</template>

<script>
import {is_login, logout} from "@/api/session";

export default {
  name: "Index",
  computed: {
    user_info() {
      return this.$store.state.user_info
    },
    active_name() {
      return this.$store.state.menu_active_name
    },
  },
  mounted() {
    is_login().then(res => {
      if (res.code !== 200) {
        this.$router.push("/login")
      }
      this.$store.commit('setUserInfo', res.data)
    })
  },
  methods: {
    handleMenuSelect(name) {
      this.$store.commit('setMenuActiveName', name)
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
  }
}
</script>

<style scoped>
html, body {
  width: 100%;
  height: 100%;
}
</style>