<template>
  <Card shadow style="min-height: 100vh">
    <Input suffix="ios-search" v-model="search_tree_input"></Input>
    <Divider>
      服务树
      <Button type="text" icon="md-refresh" @click="refreshTree"></Button>
    </Divider>
    <Tree :data="tree" :render="renderContent" @on-select-change="onSelectChange"
          @on-toggle-expand="onToggleExpend" @on-contextmenu="handleContextMenuSelected">
      <template slot="contextMenu">
        <DropdownItem @click.native="handleDeleteTreeNode" style="color: #ed4014">删除</DropdownItem>
      </template>
    </Tree>
  </Card>
</template>

<script>
import {apiGetTreeInfo, apiDeleteTree} from '@/api/tree'

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
      context_data: null,  // 右键菜单选中的数据
    }
  },
  watch: {
    search_tree_input(val) {
      this.searchTreeAndSelect(this.tree, val)
    },
  },
  mounted() {
    this.refreshTree()
  },
  methods: {
    refreshTree() {
      apiGetTreeInfo().then(res => {
        if (res.code === 200) {
          this.tree = res.data
        }
      })
    },
    renderContent(h, {root, node, data}) {
      root
      this.$set(node.node, 'contextmenu', true)
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
        h('span', {
          style: {
            textDecoration: data.is_deleted === true ? 'line-through' : '',
            color: data.is_deleted === true ? 'gray' : ''
          }
        }, data.name + (data.is_deleted === true ? '（已删除）' : '')),
      ])
    },
    onSelectChange(nodes, current_node) {
      if (current_node.expand) {
        current_node.expand = !current_node.expand
      } else {
        current_node.expand = true
      }
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
    async handleDeleteTreeNode() {
      this.$Modal.confirm({
        title: '请确认操作',
        loading: true,
        onOk: () => {
          apiDeleteTree({tree_id: this.context_data.id}).then(res => {
            if (res.code === 200) {
              this.refreshTree()
            }
            this.$Modal.remove()
          })
        }
      })
    },
    handleContextMenuSelected(data) {
      this.context_data = data
    }
  }
}
</script>

<style scoped>

</style>