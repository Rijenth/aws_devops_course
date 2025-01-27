<script setup lang="ts">
const props = defineProps<{
  message: string;
}>();
</script>

<template>
  <div class="error">{{ props.message }}</div>
</template>

<style scoped>
.error {
  color: red;
  background-color: #fdd;
  border: 1px solid #f00;
  padding: 10px;
  margin-bottom: 20px;
}
</style>
