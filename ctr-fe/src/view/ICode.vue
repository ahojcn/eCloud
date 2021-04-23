<template>
  <div>
    <Row :gutter="18">
      <Col span="5">
        <Card title="开发机列表" :padding="0" style="min-height: 100vh">
          <CellGroup @on-click="onCellClick">
            <Cell v-for="(item, index) in icode_list" :key="index" :title="item.name"
                  :to="'/icode/detail?id='+item.id" :name="item.id" :selected="selected_cell===item.id">
              <template slot="label">
                创建于
                <Time :time="item.create_time"></Time>
              </template>
            </Cell>
          </CellGroup>
        </Card>
      </Col>
      <Col span="19">
        <router-view @refreshICodeList="refreshICodeList"></router-view>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiGetICodeList} from "@/api/icode";

export default {
  name: "ICode",
  data() {
    return {
      icode_list: [],
      selected_cell: ''
    }
  },
  mounted() {
    this.refreshICodeList()
  },
  methods: {
    onCellClick(name) {
      this.selected_cell = name
    },
    refreshICodeList() {
      this.$Spin.show()
      apiGetICodeList().then(res => {
        if (res.code === 200) {
          this.icode_list = res.data
        }
        this.$Spin.hide()
      })
    },
  },
}
</script>

<style scoped>

</style>