<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { User } from "../generated/user";
import { ResponsiveTableType } from "../types/responsive_table.type";
import ResponsiveTable from "../components/ResponsiveTable.vue";
import UserApi from "../api/user.api";
import UserDetailsModal from "../widgets/UserDetailsModal.vue";
import { useToast } from "../state/useToast";

const toast = useToast();
const userApi = new UserApi();
const users = ref<User[]>([]);
const showModal = ref<boolean>(false);
const selectedUser = ref<User | null>(null);

const fetchAllUsers = async () => {
  const response = await userApi.index();

  if (typeof response === "string") {
    toast.notify(response, "error");
    return;
  }

  if (response?.users) {
    response?.users.forEach((user: User) => {
      users.value.push(user);
    });

    return;
  }

  toast.notify("Impossible de récupérer la liste des utilisateurs", "error");
};

const handleResponsiveTableRowClick = (row: Record<string, string>) => {
  if (!row.username || !row.email_address) {
    return;
  }

  const user = users.value.find(
    (user) =>
      user.username === row.username && user.email === row.email_address,
  );

  if (!user) {
    return;
  }

  selectedUser.value = user;
  showModal.value = true;
};

const tableData = computed(() => {
  let table = ref<ResponsiveTableType>({
    bold: true,
    matrix: [
      [
        { label: "username", value: "Nom d'utilisateur" },
        { label: "email_address", value: "Adresse mail" },
        { label: "first_name", value: "Prénom" },
        { label: "last_name", value: "Nom" },
        { label: "bio", value: "Biographie" },
      ],
    ],
  });

  users.value.forEach((user: User) => {
    table.value.matrix.push([
      { label: "username", value: user.username },
      { label: "email_address", value: user.email },
      { label: "first_name", value: user.firstName },
      { label: "last_name", value: user.lastName },
      { label: "bio", value: user.bio },
    ]);
  });

  return table;
});

onMounted(() => {
  fetchAllUsers();
});
</script>

<template>
  <div class="container">
    <ResponsiveTable
      :table="tableData.value"
      @responsive-table:row-click="handleResponsiveTableRowClick"
    />

    <UserDetailsModal
      :selectedUser="selectedUser"
      :showModal="showModal"
      @selected-user:deleted="
        (newSelectedUser: User) => {
          selectedUser = null;
          showModal = false;
          users = users.filter((user) => user.id !== newSelectedUser.id);
        }
      "
      @show-modal:update="
        () => {
          showModal = false;
        }
      "
    />
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  padding-top: 20px;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
}

.user-table th,
.user-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.user-table th {
  background-color: #f4f4f4;
  font-weight: bold;
}

.user-table tr:nth-child(even) {
  background-color: #f9f9f9;
}

.user-table tr:hover {
  background-color: #f1f1f1;
  cursor: pointer;
}

@media (max-width: 600px) {
  .user-table {
    display: block;
    width: 100%;
    overflow-x: auto;
    white-space: nowrap;
  }

  .user-table th,
  .user-table td {
    padding: 10px;
    font-size: 14px;
  }
}
</style>
