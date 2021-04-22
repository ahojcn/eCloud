<template>
  <Layout>
    <Header :style="{position: 'fixed', width: '100%', zIndex: 99}">
      <Menu :active-name="active_name" mode="horizontal" theme="dark" @on-select="handleMenuSelect">
        <div>
          <MenuItem name="service" to="service">
            <Icon type="md-infinite"/>
            服务
          </MenuItem>
          <MenuItem name="monitor">
            <Icon type="ios-stats"/>
            监控
          </MenuItem>
          <MenuItem name="deploy">
            <Icon type="md-cloud-upload"/>
            部署
          </MenuItem>
          <MenuItem name="resource" to="resource">
            <Icon type="md-code-working"/>
            资源
          </MenuItem>
          <Submenu name="dev">
            <template slot="title">
              <Icon type="ios-bug"/>
              开发
            </template>
            <MenuGroup title="开发机">
              <MenuItem name="create_icode" to="icode">申请开发机</MenuItem>
              <MenuItem name="icode_list" to="icode">我的开发机</MenuItem>
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
html,body {
  width: 100%;
  height: 100%;
}

.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  display: flex;
}

.layout-logo {
  width: 100px;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
}

.layout-nav {
  width: 420px;
  margin: 0 auto;
  margin-right: 20px;
}
</style>