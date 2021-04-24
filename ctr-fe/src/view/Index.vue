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
              <MenuItem name="service" :to="{name:'Service'}">
                <Icon type="md-apps"/>
                服务树信息
              </MenuItem>
              <MenuItem name="service_create" :to="{name:'TreeNodeCreate'}">
                <Icon type="md-add"/>
                创建树节点
              </MenuItem>
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
          <Submenu name="resource">
            <template slot="title">
              <Icon type="md-code-working"/>
              资源
            </template>
            <MenuGroup title="主机">
              <MenuItem name="HostDetail" :to="{name: 'HostDetail', query: {id: $route.query.id}}">
                <Icon type="md-apps"/>
                主机详情
              </MenuItem>
              <MenuItem name="HostAdd" :to="{name: 'HostAdd'}">
                <Icon type="md-add"/>
                添加主机
              </MenuItem>
              <MenuItem name="HostMonitor" :to="{name: 'HostMonitor', query: {id: $route.query.id}}">
                <Icon type="md-trending-up"/>
                主机监控
              </MenuItem>
              <MenuItem name="HostRunCmd" :to="{name: 'HostRunCmd', query: {id: $route.query.id}}">
                <Icon type="md-code-working"/>
                终端连接
              </MenuItem>
            </MenuGroup>
          </Submenu>
          <Submenu name="icode">
            <template slot="title">
              <Icon type="ios-bug"/>
              开发
            </template>
            <MenuGroup title="开发机">
              <MenuItem name="icode_list" :to="{name:'ICode'}">
                <Icon type="md-finger-print"/>
                我的开发机
              </MenuItem>
              <MenuItem name="icdoe_create" :to="{name:'ICodeCreate'}">
                <Icon type="md-add"/>
                申请开发机
              </MenuItem>
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
import {setWaterMark} from '@/util/watermask';

import {apiIsLogin, apiLogout} from "@/api/session";

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
    apiIsLogin().then(res => {
      if (res.code !== 200) {
        this.$router.push("/login")
      }
      this.$store.commit('setUserInfo', res.data)
      setWaterMark(res.data.username, res.data.email)
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
      apiLogout().then(res => {
        if (res.code === 200) {
          this.$router.go(0)
        }
      })
    },
  },
  data() {
    return {
      live2dw: {
        state: false,
        cdnPath: "/",
        type: "rem",
        position: "right",
        width: 250,
        height: 500,
        hOffset: 0,
        vOffset: -110
      },
    }
  }
}
</script>

<style scoped>
html, body {
  width: 100%;
  height: 100%;
}
</style>