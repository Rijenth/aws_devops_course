<script setup lang="ts">
import { defineProps, ref, watch } from "vue";
import ErrorMessage from "./ErrorMessage.vue";
import useInputValidate from "../composables/inputValidate";
import { InputRules } from "../types/input_rules.type";

const props = defineProps<{
  params: {
    title: string;
    value: string;
    updateable?: boolean;
  };
  placeholder?: string;
  inputType?: string;
  inputRules?: InputRules[];
  informationBubble?: boolean;
  informationMessage?: string;
}>();
const errorMessage = ref<string>("");
const localValue = ref(props.params.value);

watch(
  () => props.params.value,
  (newValue) => {
    localValue.value = newValue;
  },
);

const emit = defineEmits<{
  (event: "updated", value: string): void;
  (event: "rulesNotMet", value: boolean): void;
}>();

watch(
  () => localValue.value,
  (newValue) => {
    validate();
    emit("updated", newValue);
  },
);

const validate = () => {
  errorMessage.value = "";

  if (!props.inputRules) {
    return;
  }

  const message = useInputValidate(localValue.value, props.inputRules);

  if (!message) {
    emit("rulesNotMet", false);
    return;
  }

  emit("rulesNotMet", true);

  errorMessage.value = message;
};
</script>

<template>
  <div>
    <strong>{{ params.title }} :</strong>
    <p class="info" v-if="informationBubble && params.updateable">
      {{ informationMessage }}
    </p>
    <p v-if="!params.updateable">{{ placeholder || localValue }}</p>

    <input v-else :type="inputType || 'text'" v-model="localValue" />

    <ErrorMessage
      v-if="errorMessage && params.updateable"
      :message="errorMessage"
    />
  </div>
</template>

<style scoped>
input {
  width: 100%;
  padding: 10px;
  margin: 5px 0;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.info {
  font-size: 0.8em;
  background-color: #f9e79f;
  color: #9c8112;
  border: 1px solid #f1c40f;
  border-radius: 5px;
  padding: 10px;
}
</style>
