<template>
  <Select v-model="selected" :loading="loading" filterable :clearable="true" :remote-method="remoteSearch"
          placeholder="根据用户名或邮箱搜索" @on-change="onChange">
    <Option v-for="(user, index) in options" :key="index" :value="user.id">
      <Tag :color="user.is_active === '未激活' ? 'red' : 'green'">{{ user.is_active }}</Tag>
      {{ user.username }} - {{ user.email }}
    </Option>
  </Select>
</template>

<script>
import {apiGetUserListByUsername} from '@/api/user'

export default {
  name: "SelectUserByUsernameOrEmail",
  methods: {
    remoteSearch(query) {
      query = query.trim()
      if (query !== '') {
        this.loading = true
        apiGetUserListByUsername({username: query}).then(res => {
          this.loading = false
          if (res.code === 200) {
            this.options = res.data
          }
        })
      } else {
        this.options = []
      }
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
  }
}
</script>

<style scoped>

</style>