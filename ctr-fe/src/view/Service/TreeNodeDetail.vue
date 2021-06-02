<template>
  <Card v-if="show_tree_node_detail" shadow style="min-height: 100vh">
    <div slot="title">
      <Tag :color="$store.state.tree_node_type[tree_node_detail.user_tree_info.type].color">
        {{ $store.state.tree_node_type[tree_node_detail.user_tree_info.type].title }}
      </Tag>
      <strong style="font-size: 20px">UN={{ tree_node_detail.user_tree_info.un }}</strong>
    </div>

    <Row :gutter="18">
      <Col span="12">
        <Form :label-width="80">
          <FormItem label="描述">
            <Input readonly type="textarea" v-model="tree_node_detail.user_tree_info.description"></Input>
          </FormItem>
          <FormItem label="我的权限">
            <Input readonly v-model="tree_node_detail.user_tree_info.right_msg"></Input>
          </FormItem>
          <FormItem label="时间">
            创建于
            <Time :time="tree_node_detail.user_tree_info.create_time"></Time>
            ，最后修改于
            <Time :time="tree_node_detail.user_tree_info.update_time"></Time>
          </FormItem>
          <FormItem label="其他">
            <Input readonly type="textarea" v-model="tree_node_detail.user_tree_info.extra"></Input>
          </FormItem>
        </Form>
      </Col>
      <Col span="12">
        <Form :label-width="80">
          <FormItem label="用户组">
            <UserTreeAdd v-if="tree_node_detail.user_tree_info.rights >= 6"
                         :tree_node="tree_node_detail.user_tree_info"
                         @onAddUserTreeSuccess="onAddUserTreeSuccess">
            </UserTreeAdd>
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
          </FormItem>
        </Form>
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
        okText: '确认',
        onOk: () => {
          apiDeleteUserTree({user_id: node.user_info.id, tree_id: node.id}).then(res => {
            if (res.code === 200) {
              this.$refs.ServiceTree.refreshTree()
            }
            this.$Modal.remove()
          })
        },
        onCancel: () => {
          this.$Message.info('已取消')
        }
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