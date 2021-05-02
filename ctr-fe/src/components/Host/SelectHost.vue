<template>
  <div>
    <Select v-model="selected" @on-change="onChange">
      <Option v-for="(host_info, index) in host_list" :key="index" :value="host_info.id">
        {{ host_info.username }}@{{ host_info.ip }}:{{ host_info.port }}
        <span style="color: aqua">
          {{ host_info.description }}
        </span>
      </Option>
    </Select>
  </div>
</template>

<script>
import {apiGetHostList} from '@/api/host';

export default {
  name: "SelectHost",
  data() {
    return {
      host_list: [],
      selected: -1
    }
  },
  mounted() {
    apiGetHostList().then(res => {
      if (res.code === 200) {
        this.host_list = res.data
      }
    })
  },
  methods: {
    onChange(host_id) {
      this.$emit('onHostSelected', host_id)
    }
  }
}
</script>

<style scoped>

</style>