<script setup lang="ts">
// crédits to Nathan Shumate: https://codepen.io/nzshumate/pen/YRZPqW
import {
  ResponsiveTableMatrix,
  ResponsiveTableType,
} from "../types/responsive_table.type";

const props = defineProps<{
  table: ResponsiveTableType;
}>();

const emit = defineEmits<{
  (
    event: "responsive-table:row-click",
    rowValues: Record<string, string>,
  ): void;
}>();

const handleRowClick = (row: ResponsiveTableMatrix[]) => {
  let rowValues = row.reduce(
    (result, cell) => {
      result[cell.label] = cell.value;

      return result;
    },
    {} as Record<string, string>,
  );

  emit("responsive-table:row-click", rowValues);
};
</script>

<template>
  <div id="table-demo" class="container">
    <table :class="{ 'with-header': props.table.bold }">
      <tr
        v-for="(row, x) in props.table.matrix"
        :key="x"
        :class="{ bold: props.table.bold && x === 0 }"
        @click="handleRowClick(row)"
      >
        <td
          v-for="(col, y) in row"
          :key="y"
          :data-th="props.table.matrix[0][y].value"
        >
          {{ col.value }}
        </td>
      </tr>
    </table>
  </div>
</template>

<style scoped lang="scss">
$breakpoint: 600px;

#table-demo {
  margin: 20px 0;

  table {
    min-width: 300px;
    width: 100%;
    border-collapse: collapse;
    border: 1px solid #ddd;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);

    tr {
      cursor: pointer;
      transition:
        background-color 0.3s ease,
        transform 0.1s ease;

      &:hover {
        background-color: rgba(0, 0, 0, 0.05);
      }

      &:active {
        background-color: rgba(0, 0, 0, 0.1);
        transform: scale(0.98);
      }

      td {
        padding: 12px 15px;
        border: 1px solid #eee;
        text-align: left;
        font-size: 16px;
      }
    }

    .bold {
      background: #f8f8f8;
      font-weight: bold;

      td {
        color: #333;
        font-size: 18px;
      }
    }
  }

  .with-header {
    tr:first-of-type {
      color: #fff;
      font-weight: bold;
      text-transform: uppercase;

      td {
        padding: 15px;
        border: 1px solid #7289da;
        color: #7289da;
      }
    }

    @media (max-width: $breakpoint) {
      tr:first-of-type {
        display: none;
      }

      tr {
        display: block;
        border: 1px solid #ddd;
        margin-bottom: 10px;

        td {
          display: block;
          width: 100%;
          box-sizing: border-box;
          text-align: left;
          padding: 8px;
          border: none;
          position: relative;

          &:before {
            content: attr(data-th) ": ";
            font-weight: bold;
            display: inline-block;
            color: #666;
            margin-right: 5px;
          }
        }
      }
    }
  }
}
</style>
