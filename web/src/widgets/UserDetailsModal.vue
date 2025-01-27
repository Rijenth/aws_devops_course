<script setup lang="ts">
import { ref } from "vue";
import BasicModal from "../components/BasicModal.vue";
import UserApi from "../api/user.api";
import { User } from "../generated/user";
import OnclickButton from "../components/OnclickButton.vue";
import UserPersonalInfo from "./UserPersonalInfo.vue";

const props = defineProps<{
  selectedUser: User | null;
  showModal: boolean;
}>();

const emit = defineEmits<{
  (event: "selected-user:deleted", user: User): void;
  (event: "show-modal:update", value: boolean): void;
}>();

const userApi = new UserApi();
const errorMessage = ref<string | null>(null);

const handleDeleteUserFromList = async () => {
  errorMessage.value = null;

  if (!props.selectedUser) {
    return;
  }

  const response = await userApi.delete(props.selectedUser.id);

  if (typeof response === "string") {
    errorMessage.value = response;
    return;
  }

  if (response?.success) {
    emit("selected-user:deleted", props.selectedUser);
    return;
  }

  errorMessage.value =
    "Une erreur est survenue lors de la suppression de l'utilisateur";
};
</script>

<template>
  <BasicModal
    :show="props.showModal"
    @basic-modal:close="emit('show-modal:update', false)"
  >
    <div class="modal-children">
      <h2 class="title">
        Fiche de
        {{
          selectedUser?.lastName.toUpperCase() + " " + selectedUser?.firstName
        }}
      </h2>

      <hr class="separator" />

      <UserPersonalInfo :username="props.selectedUser?.username" />

      <div class="modal-actions">
        <OnclickButton
          text="Fermer"
          :onButtonClick="
            () => {
              emit('show-modal:update', false);
            }
          "
        />

        <OnclickButton
          text="Retirer de votre liste"
          :onButtonClick="handleDeleteUserFromList"
          backgroundColor="#d93939"
          hoverColor="#b72f2f"
        />
      </div>
    </div>
  </BasicModal>
</template>

<style scoped>
.title {
  text-align: center;
  border: 1px solid #7289da;
  border-radius: 5px;
  padding: 10px;
}

.separator {
  margin: 20px 0;
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
</style>
