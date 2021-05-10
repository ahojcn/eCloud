<template>
  <div>
    <Button type="primary" ghost size="small" icon="md-add" @click="show_form_modal = true">
      添加用户
    </Button>

    <Modal transfer v-model="show_form_modal" title="服务树添加用户">
      <Form :label-width="80">
        <FormItem label="当前节点">
          <Tag :color="$store.state.tree_node_type[tree_node.type].color">
            {{ $store.state.tree_node_type[tree_node.type].title }}
          </Tag>
          {{ tree_node.name }}
        </FormItem>
        <FormItem label="用户">
          <SelectUserByUsernameOrEmail @onChange="onUserSelected"></SelectUserByUsernameOrEmail>
        </FormItem>
        <FormItem label="权限">
          <Select v-model="form.rights">
            <Option v-for="(item, index) in $store.state.user_tree_rights" :key="index" :value="item.value">
              {{ item.label }}
            </Option>
          </Select>
        </FormItem>
        <FormItem>
          <Alert type="error">权限具有继承关系，请谨慎操作！</Alert>
        </FormItem>
      </Form>

      <div slot="footer">
        <Button type="primary" long @click="handleUserTreeAddClick">添加</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import {apiAddUserTree} from '@/api/tree';
import SelectUserByUsernameOrEmail from "@/components/User/SelectUserByUsernameOrEmail";

export default {
  name: "UserTreeAdd",
  components: {SelectUserByUsernameOrEmail},
  props: {
    tree_node: {
      required: true,
      type: Object,
    },
  },
  data() {
    return {
      show_form_modal: false,
      form: {
        user_id: -1,
        tree_id: -1,
        rights: -1
      },
    }
  },
  methods: {
    onUserSelected(id) {
      this.form.user_id = id
    },
    handleUserTreeAddClick() {
      apiAddUserTree(this.form).then(res => {
        if (res.code === 200) {
          this.show_form_modal = false
          this.$emit('onAddUserTreeSuccess')
        }
      })
    },
  },
  watch: {
    tree_node(node) {
      this.form.tree_id = node.id
    }
  },
}
</script>

<style scoped>

</style>