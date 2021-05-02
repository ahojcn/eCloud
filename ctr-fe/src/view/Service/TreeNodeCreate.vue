<template>
  <div>
    <Card shadow style="min-height: 100vh">
      <p slot="title" style="text-align: center">创建节点</p>
      <Row :gutter="18">
        <Col span="12">
          <Alert type="error">
            节点的权限是具有继承关系的，建议给其他用户添加最小的节点权限。
          </Alert>
          <Alert type="warning">
            节点分级如下：<br/>
            -------------------------<br/>
            Namespace（命名空间）<br/>
            |--ProductLine（产品线）<br/>
            |--|--Group（用户组）<br/>
            |--|--|--Service（服务）<br/>
            |--|--|--|--Cluster（集群）<br/>
            -------------------------
          </Alert>
          <Alert type="info">
            每一个父级节点都可以包含 1 个或多个子节点。<br/>
            每个服务树的节点，不论是根节点还是叶子节点都会有一个唯一的标识(UniqueName)。<br/>
            每个 UniqueName 都是由 '.' 分割的。<br/>
            例如：op.xiaoniu.zhieasy 代表 运维组.小牛项目.致易青年。<br/>
          </Alert>
        </Col>
        <Col span="12" :style="{textAlign: 'center'}">
          <Form :model="add_tree_node_form_data" :label-width="80">
            <FormItem prop="name" label="节点名称">
              <Input type="text" v-model="add_tree_node_form_data.name"></Input>
            </FormItem>
            <FormItem prop="description" label="节点描述">
              <Input type="textarea" v-model="add_tree_node_form_data.description"></Input>
            </FormItem>
            <FormItem prop="type" label="节点类型">
              <Select v-model="add_tree_node_form_data.type">
                <Tag slot="prefix" :color="$store.state.tree_node_type[add_tree_node_form_data.type].color">
                  {{ $store.state.tree_node_type[add_tree_node_form_data.type].title }}
                </Tag>
                <Option v-for="(item, index) in $store.state.tree_node_type_selector" :key="index"
                        :value="item.value">
                  {{ item.label }}
                </Option>
              </Select>
            </FormItem>
            <FormItem prop="parent_id" label="父节点">
              <SelectTreeNodeByName @onChange="onTreeNodeSelected"></SelectTreeNodeByName>
            </FormItem>
            <FormItem>
              <Button type="primary" long @click="handleAddTreeNodeBtnClick">添加服务树节点</Button>
            </FormItem>
          </Form>
        </Col>
      </Row>
    </Card>
  </div>
</template>

<script>
import {apiAddTreeNode} from "@/api/tree";

import SelectTreeNodeByName from "@/components/Tree/SelectTreeNodeByName";

export default {
  name: "TreeNodeCreate",
  components: {SelectTreeNodeByName},
  data() {
    return {
      add_tree_node_form_data: {
        name: '',
        description: '',
        type: 0,
        parent_id: 0
      },
    }
  },
  methods: {
    handleAddTreeNodeBtnClick() {
      apiAddTreeNode(this.add_tree_node_form_data).then(res => {
        if (res.code === 200) {
          this.$emit('onTreeNodeAddSuccessful')
        }
      })
    },
    onTreeNodeSelected(id) {
      this.add_tree_node_form_data.parent_id = id
    },
  },
}
</script>

<style scoped>

</style>