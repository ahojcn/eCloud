<template>
  <div>
    <VueShell v-if="!loading" :banner="shell_banner" @shell_output="onRunCommand"
              :shell_input="send_to_terminal"></VueShell>
  </div>
</template>

<script>
import {apiGetHostList, apiRunCommand} from "@/api/host";

import VueShell from 'vue-shell';

export default {
  name: "HostRunCmd",
  components: {VueShell},
  data() {
    return {
      host_info: {},
      send_to_terminal: '',
      loading: true
    }
  },
  computed: {
    shell_banner: function () {
      return {
        header: this.host_info.ip,
        subHeader: "数据无价，谨慎操作！",
        helpHeader: '你的所有操作都将会被记录！',
        emoji: {
          first: "e-Cloud 谨慎操作",
          second: "e-cloud 数据无价",
          time: 750
        },
        sign: "$",
        img: {
          align: "right",
          link: "https://ae01.alicdn.com/kf/U11c312b4a0034a338e56760694ad7906J.jpg",
          width: 200,
          height: 200
        }
      }
    },
  },
  methods: {
    onRunCommand(val) {
      this.$Spin.show()
      apiRunCommand({
        host_id: this.host_info.id,
        cmd: val
      }).then(res => {
        if (res.code === 200) {
          this.send_to_terminal = res.data
        } else {
          this.send_to_terminal = res.msg
        }
        this.$Spin.hide()
      })
    },
    getHostInfo(id) {
      apiGetHostList({"id": id}).then(res => {
        if (res.code === 200) {
          this.host_info = res.data[0]
          this.loading = false
        }
      })
    },
  },
  mounted() {
    this.getHostInfo({id: this.$route.query.id})
  }
}
</script>

<style scoped>

</style>