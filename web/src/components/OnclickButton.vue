<script setup lang="ts">
import { defineProps } from "vue";

const props = defineProps<{
  onButtonClick: () => void;
  text: string;
  backgroundColor?: string;
  hoverColor?: string;
  disabled?: boolean;
}>();

const handleClick = () => {
  props.onButtonClick();
};
</script>

<template>
  <button
    class="btn"
    @click="handleClick"
    :disabled="props.disabled"
    :style="{
      '--bg-color': props.backgroundColor || '#7289da',
      '--hover-color': props.hoverColor || '#5d6eab',
      cursor: props.disabled ? 'not-allowed' : 'pointer',
    }"
  >
    <span>{{ props.text }}</span>
  </button>
</template>

<style scoped>
.btn {
  margin: 10px 0;
  display: inline-block;
  padding: 10px;
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
  background-color: var(--bg-color);
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition:
    background-color 0.3s ease,
    transform 0.1s ease;
}

.btn:hover {
  background-color: var(--hover-color);
}

.btn:active {
  background-color: darken(var(--hover-color), 10%);
  transform: scale(0.98);
}
</style>
