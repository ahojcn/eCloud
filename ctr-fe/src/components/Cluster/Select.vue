<template>
  <div>
    <Select v-model="selected" filterable @on-change="onSelectChange">
      <Option v-for="item in d" :key="item.id" :value="item.id">
        {{ item.name }}
      </Option>
    </Select>
  </div>
</template>

<script>
import {apiClusterList} from "@/api/cluster";

export default {
  name: "ClusterSelect",
  props: {
    tree_id: {
      require: true,
      type: Number
    },
  },
  data() {
    return {
      d: [],
      selected: '',
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    getData() {
      apiClusterList({tree_id: this.tree_id}).then(res => {
        if (res.code === 200) {
          this.d = res.data
        }
      })
    },
    onSelectChange(value) {
      this.$emit('onClusterSelected', value)
    },
  }
}
</script>

<style scoped>

</style>