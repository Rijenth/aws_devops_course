<script setup lang="ts">
import { defineProps, ref, watch } from "vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { Timestamp } from "../generated/google/protobuf/timestamp";

const props = defineProps<{
  title: string;
  timestamp: Timestamp | undefined;
  updateable?: boolean;
}>();

const locale = ref<string>("fr");
const dateValue = ref<Date | null>(convertTimestampToDate(props.timestamp));

function convertTimestampToDate(timestamp?: Timestamp): Date | null {
  if (!timestamp) return null;
  return new Date(Number(timestamp.seconds) * 1000);
}

function convertDateToTimestamp(date: Date): Timestamp {
  return {
    seconds: BigInt(Math.floor(date.getTime() / 1000)),
    nanos: (date.getTime() % 1000) * 1e6,
  };
}

watch(
  () => props.timestamp,
  (newTimestamp) => {
    let newDate = convertTimestampToDate(newTimestamp);

    if (dateValue.value?.getTime() !== newDate?.getTime()) {
      dateValue.value = convertTimestampToDate(newTimestamp);
    }
  },
  { immediate: true },
);

const emit = defineEmits<{
  (event: "updated", timestamp: Timestamp): void;
}>();

watch(dateValue, (newDate) => {
  if (newDate) {
    emit("updated", convertDateToTimestamp(newDate));
  }
});
</script>

<template>
  <div>
    <strong>{{ title }} :</strong>
    <Datepicker
      v-if="updateable"
      v-model="dateValue"
      :format="
        (date?: Date) =>
          date
            ? date.toLocaleDateString('fr-FR')
            : new Date().toLocaleDateString('fr-FR')
      "
      :clearable="false"
      :enableTimePicker="false"
      :autoApply="true"
      :maxDate="new Date()"
      :locale="locale"
    />
    <p v-else>{{ dateValue ? dateValue.toLocaleDateString("fr-FR") : "" }}</p>
  </div>
</template>
