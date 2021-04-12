<template>
  <div>
    <transition>
      <img class="animate__animated animate__flip"
           :style="{width: size}" id="logo"
           src="@/assets/logo.png" alt="logo" @click="$router.push('/')">
    </transition>
  </div>
</template>

<script>
export default {
  name: "Logo",
  props: {
    size: {
      type: String,
      default: '100px'
    }
  }
}
</script>

<style scoped>
</style>