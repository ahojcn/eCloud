<template>
  <div id="register">
    <Logo size="10%"></Logo>
    <div style="padding-top: 30px"></div>
    <Input v-model="form.username" placeholder="用户名" style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Input v-model="form.password" placeholder="密码" type="password" password style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Input v-model="form.r_password" placeholder="重复密码" type="password" password style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Input v-model="form.email" placeholder="邮箱" password style="width: 300px"/>
    <div style="padding-top: 10px"></div>
    <Button ghost type="primary" to="/login" style="width: 100px">登录</Button>
    <Button type="primary" style="width: 200px" @click="handleRegisterBtnClick">注册</Button>
  </div>
</template>

<script>
import Logo from "@/components/Logo";
import {register} from '@/api/user'
import {apiIsLogin} from "@/api/session";

export default {
  name: "Register",
  components: {Logo},
  data() {
    return {
      form: {
        username: "",
        password: "",
        r_password: "",
        email: ""
      },
    }
  },
  mounted() {
    apiIsLogin().then(res => {
      if (res.code === 200) {
        this.$router.push("/")
      }
    })
  },
  methods: {
    handleRegisterBtnClick() {
      register(this.form).then(res => {
        console.log(res)
        if (res.code === 200) {
          alert("ok")
        }
      })
    }
  }
}
</script>

<style scoped>
#register {
  text-align: center;
  margin-top: 60px;
}
</style>