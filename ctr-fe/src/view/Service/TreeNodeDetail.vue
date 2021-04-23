<template>
  <Card v-if="show_tree_node_detail" shadow style="min-height: 100vh">
    <div slot="title" style="text-align: center">
      <Tag size="large" :color="$store.state.tree_node_type[tree_node_detail.user_tree_info.type].color">
        {{ $store.state.tree_node_type[tree_node_detail.user_tree_info.type].title }}
      </Tag>
      <h1>{{ tree_node_detail.user_tree_info.un }}</h1>
    </div>

    <Row :gutter="18">
      <Col span="12">
        <Card shadow style="min-height: 100vh">
          <p slot="title" style="text-align: center">服务树信息</p>
          <p>详细描述：{{ tree_node_detail.user_tree_info.description }}</p>
          <p>我的权限：{{ tree_node_detail.user_tree_info.right_msg }}</p>
          <p>创建时间：{{ tree_node_detail.user_tree_info.create_time }}</p>
          <p>更新时间：{{ tree_node_detail.user_tree_info.update_time }}</p>
          <p>其他信息：{{ tree_node_detail.user_tree_info.extra }}</p>
          <p>子节点：{{ tree_node_detail.children.length }}</p>
        </Card>
      </Col>
      <Col span="12">
        <Card shadow>
          <p slot="title" style="text-align: center">用户组信息</p>
          <Tooltip transfer v-for="(user_info, index) in tree_node_detail.users" :key="index">
            <Tag v-if="tree_node_detail.user_tree_info.user_info.id === user_info.user_info.id"
                 color="red">{{ user_info.user_info.username }}（管理员）
            </Tag>
            <Tag v-else>{{ user_info.user_info.username }}</Tag>
            <div slot="content">
              <p>权限：{{ user_info.right_msg }}</p>
              <p>邮箱：{{ user_info.user_info.email }}</p>
              <Button size="small" type="error"
                      @click="handleDeleteUserTree(user_info)">删除
              </Button>
            </div>
          </Tooltip>
          <UserTreeAdd v-if="tree_node_detail.user_tree_info.rights >= 6"
                       :tree_node="tree_node_detail.user_tree_info"
                       @onAddUserTreeSuccess="onAddUserTreeSuccess"
          ></UserTreeAdd>
        </Card>
      </Col>
    </Row>
  </Card>
</template>

<script>
import {apiDeleteUserTree, apiGetTreeInfo} from '@/api/tree';
import UserTreeAdd from "@/view/Service/UserTreeAdd";

export default {
  name: "TreeNodeDetail",
  components: {UserTreeAdd},
  computed: {
    id: function () {
      return this.$route.query.id
    }
  },
  watch: {
    id(new_id) {
      this.refreshTreeNodeInfo(new_id)
    }
  },
  data() {
    return {
      tree_node_detail: {},
      show_tree_node_detail: false,
    }
  },
  methods: {
    async handleDeleteUserTree(node) {
      this.$Modal.confirm({
        title: '请确认删除操作',
        loading: true,
        okText: '已确认',
        onOk: () => {
          apiDeleteUserTree({user_id: node.user_info.id, tree_id: node.id}).then(res => {
            if (res.code === 200) {
              this.$refs.ServiceTree.refreshTree()
            }
            this.$Modal.remove()
          })
        },
      })
    },
    onAddUserTreeSuccess() {
      this.$emit('onAddUserTreeSuccess')
    },
    refreshTreeNodeInfo(id) {
      apiGetTreeInfo({"id": id}).then(res => {
        if (res.code === 200) {
          this.tree_node_detail = res.data
          this.show_tree_node_detail = true
        }
      })
    }
  },
  mounted() {
    this.refreshTreeNodeInfo(this.id)
  }
}
</script>

<style scoped>

</style>