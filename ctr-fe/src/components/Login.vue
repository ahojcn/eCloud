<template>
  <div id="login">
    <div style="text-align:center">
      <Logo size="10%"></Logo>

      <div style="padding-top: 30px"></div>
      <Input v-model="login_form.username" placeholder="用户名" style="width: 300px"/>
      <div style="padding-top: 10px"></div>
      <Input v-model="login_form.password" placeholder="密码" type="password" password style="width: 300px"/>
      <div style="padding-top: 10px"></div>
      <Button type="text">忘记密码？</Button>
      <Button type="primary" style="width: 200px" @click="handleLoginClick">登录</Button>
    </div>
  </div>
</template>

<script>
import Logo from "@/components/Logo";
import {login, getYuQueOAuth} from "@/api/session"

export default {
  name: "Login",
  components: {Logo},
  data() {
    return {
      login_form: {
        username: "",
        password: ""
      }
    }
  },
  methods: {
    handleLoginClick() {
      login(this.login_form).then(res => {
        console.log(res)
        getYuQueOAuth({"username": "a"}).then(res => {
          console.log(res)
        })
      })
    },
  }
}
</script>

<style scoped>
#login {
  text-align: center;
  margin-top: 60px;
}
</style>