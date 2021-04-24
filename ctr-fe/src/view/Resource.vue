<template>
  <div>
    <Row :gutter="18">
      <Col span="4">
        <Card title="主机列表" :padding="0" style="min-height: 100vh">
          <CellGroup @on-click="onCellClick">
            <Cell v-for="(host_info, index) in host_list" :key="index" :title="host_info.ip"
                  :to="'?id='+host_info.id" :name="host_info.id"
                  :selected="selected_cell===host_info.id">
              <template slot="label">
                创建于
                <Time :time="host_info.create_time"></Time>
              </template>
            </Cell>
          </CellGroup>
        </Card>
      </Col>
      <Col span="20">
        <router-view @onAddHostUserOk="refreshHostList"></router-view>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiGetHostList} from "@/api/host";

export default {
  name: "Resource",
  props: {},
  data() {
    return {
      selected_cell: 0,
      host_list: [],
      host_list_loading: true,
    }
  },
  methods: {
    onCellClick(name) {
      this.selected_cell = name
    },
    refreshHostList() {
      apiGetHostList().then(res => {
        if (res.code === 200) {
          this.host_list = res.data
          this.host_list_loading = false
        }
      })
    },
  },
  mounted() {
    this.refreshHostList()

    let id = this.$route.query.id
    this.selected_cell = parseInt(id === '' ? '0' : id, 10)
  }
}
</script>

<style scoped>

</style>