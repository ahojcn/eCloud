<template>
  <Card>
    <Input suffix="ios-search" v-model="search_tree_input"></Input>
    <Divider>
      服务树
      <Button type="text" icon="md-refresh" @click="refreshTree"></Button>
    </Divider>
    <Tree :data="tree" :render="renderContent" @on-select-change="onSelectChange"
          @on-toggle-expand="onToggleExpend"></Tree>
  </Card>
</template>

<script>
import {apiGetTreeInfo} from '@/api/tree'

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
      search_tree_input: '',
    }
  },
  watch: {
    search_tree_input(val) {
      this.searchTreeAndSelect(this.tree, val)
    }
  },
  mounted() {
    this.refreshTree()
  },
  methods: {
    refreshTree() {
      apiGetTreeInfo().then(res => {
        if (res.code === 200) {
          this.tree = res.data
          if (this.tree.length !== 0) {
            this.tree[0].expand = true
            this.onToggleExpend(this.tree[0])
          }
        }
      })
    },
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
      current_node.expand = true
      this.$emit('onTreeNodeSelected', current_node)
    },
    onToggleExpend(current_node) {
      this.$emit('onToggleExpend', current_node)
    },
    searchTreeAndSelect(tree_node, val) {
      if (tree_node === [] || tree_node === undefined) {
        return undefined
      } else {
        for (let i = 0; i < tree_node.length; i++) {
          if (tree_node[i].name === val || tree_node[i].un === val
              || tree_node[i].name.startsWith(val) || tree_node[i].un.startsWith(val)) {
            this.$set(tree_node[i], 'selected', true)
            this.$set(tree_node[i], 'expand', true)
            this.onToggleExpend(tree_node[i])
            return i
          } else {
            if (this.searchTreeAndSelect(tree_node[i].children, val) !== undefined) {
              this.$set(tree_node[i], 'selected', true)
              this.$set(tree_node[i], 'expand', true)
              return i
            } else {
              this.$set(tree_node[i], 'selected', false)
              this.$set(tree_node[i], 'expand', false)
            }
          }
        }
      }
    },
  }
}
</script>

<style scoped>

</style>