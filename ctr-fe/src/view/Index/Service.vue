<template>
  <div>
    <Row :gutter="16">
      <Col span="4">
        <ServiceTree ref="ServiceTree" @onTreeNodeSelected="onTreeSelectedOrExpand"
                     @onToggleExpend="onTreeSelectedOrExpand"></ServiceTree>
      </Col>
      <Col span="20">
        <Card v-if="show_tree_node_detail">
          <div slot="title">
            <Tag :color="$store.state.tree_node_type[tree_node_detail.user_tree_info.type].color">
              {{ $store.state.tree_node_type[tree_node_detail.user_tree_info.type].title }}
            </Tag>
            <span style="font-size: 20px">{{ tree_node_detail.user_tree_info.un }}</span>
          </div>
          <div>
            <Row :gutter="18">
              <Col span="8">
                <Alert>
                  服务树信息：
                  <template slot="desc">
                    <p>详细描述：{{ tree_node_detail.user_tree_info.description }}</p>
                    <p>我的权限：{{ tree_node_detail.user_tree_info.right_msg }}</p>
                    <p>创建时间：{{ tree_node_detail.user_tree_info.create_time }}</p>
                    <p>更新时间：{{ tree_node_detail.user_tree_info.update_time }}</p>
                    <p>其他信息：{{ tree_node_detail.user_tree_info.extra }}</p>
                    <p>子节点：{{tree_node_detail.children.length}}</p>
                  </template>
                </Alert>
              </Col>
              <Col span="8">
                <Alert>
                  用户组信息：
                  <template slot="desc">
                    <Tooltip transfer v-for="(user_info, index) in tree_node_detail.users" :key="index">
                      <span>{{ user_info.user_info.username }},</span>
                      <div slot="content">
                        <p>权限：{{ user_info.right_msg }}</p>
                        <p>邮箱：{{ user_info.user_info.email }}</p>
                        <Button size="small" type="error"
                                @click="handleDeleteUserTree(user_info)">删除
                        </Button>
                      </div>
                    </Tooltip>
                    <AddUserTreeBtn v-if="tree_node_detail.user_tree_info.rights >= 6"
                                    :tree_node="tree_node_detail.user_tree_info"
                                    @onAddUserTreeSuccess="onAddUserTreeSuccess"
                    ></AddUserTreeBtn>
                  </template>
                </Alert>
              </Col>
              <Col span="8">
                <AddTreeNodeForm @onTreeNodeAddSuccessful="onTreeNodeAddSuccessful"></AddTreeNodeForm>
              </Col>
            </Row>
          </div>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
import {apiGetTreeInfo, apiDeleteUserTree} from '@/api/tree'
import ServiceTree from "@/components/ServiceTree";
import AddUserTreeBtn from "@/components/AddUserTreeBtn";
import AddTreeNodeForm from "@/components/AddTreeNodeForm";

export default {
  name: "Service",
  components: {AddTreeNodeForm, AddUserTreeBtn, ServiceTree},
  methods: {
    onTreeSelectedOrExpand(node) {
      apiGetTreeInfo({id: node.id}).then(res => {
        if (res.code === 200) {
          this.tree_node_detail = res.data
          this.show_tree_node_detail = true
        }
      })
    },
    onTreeNodeAddSuccessful() {
      this.$refs.ServiceTree.refreshTree()
    },
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
      this.$refs.ServiceTree.refreshTree()
    },
  },
  computed: {},
  data() {
    return {
      tree_node_detail: {},
      show_tree_node_detail: false,
      user_table_columns: [
        {title: 'id', key: 'user_info.id'},
        {title: '用户名', key: 'user_info.username'},
        {title: '邮箱', key: 'email'},
        {title: '电话', key: 'phone'},
      ],
    }
  },
}
</script>

<style scoped>

</style>