<template>
  <div>
    <Button>Default</Button>
    <Button type="primary">Primary</Button>
    <Button type="dashed">Dashed</Button>
    <Button type="text">Text</Button>
    <br><br>
    <Button type="info">Info</Button>
    <Button type="success">Success</Button>
    <Button type="warning">Warning</Button>
    <Button type="error">Error</Button>
  </div>
</template>

<script>
// import {getYuQueOAuth} from "@/api/session"

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  mounted() {
    // getYuQueOAuth({"username": "ahojcn"}).then(res => {
    //   console.log(res)
    // })
  }
}
</script>

<style scoped>
</style>
