<script setup lang="ts">
import { defineProps } from "vue";

defineProps({
  label: { type: String, required: true },
  id: { type: String, required: true },
  type: { type: String, default: "text" },
  modelValue: { type: String, default: "" },
  placeholder: { type: String, default: "" },
});

const emit = defineEmits<{
  (event: "update:modelValue", value: string): void;
}>();

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement | null;

  if (target) {
    emit("update:modelValue", target.value);
  }
};
</script>

<template>
  <div class="form-group">
    <label :for="id">{{ label }}</label>
    <input
      :id="id"
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      @input="handleInput"
    />
  </div>
</template>

<style scoped>
.form-group {
  margin-bottom: 15px;
}

label {
  font-size: 18px;
  color: #666;
}

input {
  width: 100%;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ddd;
  font-size: 14px;
  box-sizing: border-box;
}
</style>
