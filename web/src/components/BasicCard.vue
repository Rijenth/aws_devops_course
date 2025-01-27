<script setup lang="ts">
defineProps<{
  title?: string;
}>();
</script>

<template>
  <div class="container">
    <h2>{{ title }}</h2>
    <slot />
  </div>
</template>

<style scoped>
.container {
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.container h2 {
  font-size: 1.5em;
  margin-bottom: 15px;
}

@media screen and (min-width: 768px) {
  .container {
    width: 100%;
    padding: 20px;
  }
}
</style>
