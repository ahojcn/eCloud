<template>
  <div>
    <Card>
      <Tree :data="tree" :render="renderContent" @on-select-change="onSelectChange"></Tree>
    </Card>
  </div>
</template>

<script>
import {get_tree} from '@/api/tree'

export default {
  name: "ServiceTree",
  data() {
    return {
      tree: [],
      tree_node_type: {
        "4": {title: "N", color: "#f56a00"},
        "3": {title: "P", color: "#7265e6"},
        "2": {title: "G", color: "#ffbf00"},
        "1": {title: "S", color: "#87d068"},
        "0": {title: "C", color: "#00a2ae"}
      },
    }
  },
  mounted() {
    get_tree().then(res => {
      this.tree = res.data
      console.log(res)
    })
  },
  methods: {
    renderContent(h, {root, node, data}) {
      root, node
      return h('span', {
        style: {
          display: 'inline-block',
          width: '100%'
        }
      }, [
        h('Avatar', {
          props: {size: 'small'},
          style: {
            backgroundColor: this.tree_node_type[data.type].color,
            marginRight: "10px"
          }
        }, this.tree_node_type[data.type].title),
        h('span', {}, data.name)
      ])
    },
    onSelectChange(nodes, current_node) {
      this.$emit('onTreeNodeSelected', current_node)
    },
  }
}
</script>

<style scoped>

</style>