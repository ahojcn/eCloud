<template>
  <div id="login">
    <Logo size="10%"></Logo>

    <div style="padding-top: 30px"></div>
    <Input v-model="login_form.username" placeholder="用户名" style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Input v-model="login_form.password" placeholder="密码" type="password" password style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Button ghost type="primary" to="/register" style="width: 100px">注册</Button>
    <Button type="primary" style="width: 200px" @click="handleLoginClick">登录</Button>
  </div>
</template>

<script>
import Logo from "@/components/Logo";
import {is_login, login} from "@/api/session"

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
  mounted() {
    is_login().then(res => {
      if (res.code === 200) {
        this.$router.push("/")
      }
    })
  },
  methods: {
    handleLoginClick() {
      login(this.login_form).then(res => {
        if (res.code === 200) {
          this.$router.push("/")
        }
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