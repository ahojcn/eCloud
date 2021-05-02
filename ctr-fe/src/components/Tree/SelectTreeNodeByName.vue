<template>
  <Select @on-change="onChange" filterable v-model="selected" :loading="loading"
          :remote-method="remoteSearch" :clearable="true">
    <Option :value="0">
      <Tag :color="$store.state.tree_node_type[4].color">{{ $store.state.tree_node_type[4].title }}</Tag>
      根节点
    </Option>
    <Option v-for="(tree, index) in options" :key="index" :value="tree.id">
      <Tag :color="$store.state.tree_node_type[tree.type].color">{{ $store.state.tree_node_type[tree.type].title }}
      </Tag>
      {{ tree.name }} - {{ tree.description }}
    </Option>
  </Select>
</template>

<script>
import {apiGetTreeInfo} from "@/api/tree";

export default {
  name: "SelectTreeNodeByName",
  methods: {
    remoteSearch(query) {
      this.loading = true
      apiGetTreeInfo({name: query}).then(res => {
        this.loading = false
        if (res.code === 200) {
          this.options = res.data
        }
      })
    },
    onChange(id) {
      this.$emit('onChange', id)
    }
  },
  data() {
    return {
      selected: 0,
      loading: false,
      options: []
    }
  },
}
</script>

<style scoped>

</style>