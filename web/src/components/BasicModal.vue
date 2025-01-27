<script setup lang="ts">
import { defineProps, defineEmits, ref, watch } from "vue";

const props = defineProps({
  show: {
    type: Boolean,
    required: true,
  },
});

const emit = defineEmits(["basic-modal:close"]);

const visible = ref(props.show);

watch(
  () => props.show,
  (newValue) => {
    visible.value = newValue;
  },
);

const closeModal = () => {
  emit("basic-modal:close");
  visible.value = false;
};
</script>

<template>
  <div v-if="visible" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-children">
        <slot />
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  max-width: 500px;
  width: 100%;
  animation: modal-appear 0.3s ease-out;
}

.modal-children {
  padding: 20px;
}

@keyframes modal-appear {
  from {
    transform: translateY(-100px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>
