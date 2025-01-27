<script setup lang="ts">
import MainNavbar from "../widgets/MainNavbar.vue";
</script>

<template>
  <div class="container">
    <MainNavbar />

    <div class="main-body">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.container {
  height: 100vh;
}

.main-body {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-grow: 1;
  padding: 20px;
}

.main-body p {
  margin: 0;
  font-size: 18px;
  text-align: center;
}
</style>
