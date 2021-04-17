<template>
  <Card>
    <p slot="title">添加节点</p>
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
  </Card>
</template>

<script>
import SelectTreeNodeByName from "@/components/SelectTreeNodeByName";
import {apiAddTreeNode} from "@/api/tree";

export default {
  name: "AddTreeNodeForm",
  components: {SelectTreeNodeByName},
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
  data() {
    return {
      add_tree_node_form_data: {
        name: '',
        description: '',
        type: 0,
        parent_id: 0
      },
    }
  }
}
</script>

<style scoped>

</style>