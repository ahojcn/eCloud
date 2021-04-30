<template>
  <div>
    <h1>接入层信息</h1>
  </div>
</template>

<script>
import {apiGetTreeInfo} from "@/api/tree";
import {apiGetRouterInfo} from "@/api/m_router";

export default {
  name: "Router",
  computed: {
    id: function () {
      return this.$route.query.id
    }
  },
  watch: {
    id(new_id) {
      this.refreshTreeNodeInfo(new_id)
    },
  },
  mounted() {
    this.refreshTreeNodeInfo(this.id)
    this.refreshRouterInfo({ns_id: this.id})
  },
  data() {
    return {
      tree_node_info: {},
      router_info: {},
    }
  },
  methods: {
    refreshTreeNodeInfo(id) {
      apiGetTreeInfo({"id": id}).then(res => {
        if (res.code === 200) {
          this.tree_node_info = res.data
        }
      })
    },
    refreshRouterInfo(obj) {
      apiGetRouterInfo(obj).then(res => {
        if (res.code === 200) {
          this.router_info = res.data
        }
      })
    },
  },
}
</script>

<style scoped>

</style>