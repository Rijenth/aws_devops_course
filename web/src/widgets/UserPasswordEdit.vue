<script setup lang="ts">
import { ref, watch } from "vue";
import BasicCard from "../components/BasicCard.vue";
import EditableParagraph from "../components/EditableParagraph.vue";
import OnclickButton from "../components/OnclickButton.vue";
import ErrorMessage from "../components/ErrorMessage.vue";
import AuthApi from "../api/auth.api";
import { useAuth } from "../state/useAuth";

const auth = useAuth();
const authApi = new AuthApi();

const currentPassword = ref<string>("");
const newPassword = ref<string>("");
const passwordConfirmation = ref<string>("");

const currentPasswordErrorMessage = ref<string>("");
const newPasswordErrorMessage = ref<string>("");

const updateMode = ref<boolean>(false);
const updateButtonDisabled = ref<boolean>(false);

const updateUserPassword = async () => {
  currentPasswordErrorMessage.value = "";
  newPasswordErrorMessage.value = "";

  const response = await authApi.updatePassword(
    auth.state.username ?? "",
    currentPassword.value,
    newPassword.value,
  );

  if (typeof response === "string") {
    currentPasswordErrorMessage.value = response;
    return;
  }

  if (response.success) {
    currentPassword.value = "";
    newPassword.value = "";
    passwordConfirmation.value = "";
    return;
  }
};

const checkPasswordConfirmation = () => {
  if (newPassword.value !== passwordConfirmation.value) {
    newPasswordErrorMessage.value = "Les mots de passe ne correspondent pas.";
    updateButtonDisabled.value = true;
    return;
  }

  newPasswordErrorMessage.value = "";
};

watch(
  () => [newPassword.value, passwordConfirmation.value],
  () => {
    checkPasswordConfirmation();
  },
);
</script>

<template>
  <BasicCard :title="'Changer mon mot de passe'">
    <ErrorMessage
      v-if="currentPasswordErrorMessage || newPasswordErrorMessage"
      :message="currentPasswordErrorMessage || newPasswordErrorMessage"
    />

    <EditableParagraph
      :params="{
        title: 'Mot de passe actuel',
        value: currentPassword,
        updateable: updateMode,
      }"
      :placeholder="'*******'"
      :inputType="'password'"
      :inputRules="['required', { minLength: 5 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          currentPassword = value;
        }
      "
    />

    <EditableParagraph
      v-if="updateMode"
      :params="{
        title: 'Nouveau mot de passe',
        value: '',
        updateable: updateMode,
      }"
      :inputType="'password'"
      :inputRules="[{ minLength: 5 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          newPassword = value;
        }
      "
    />

    <EditableParagraph
      v-if="updateMode"
      :params="{
        title: 'Confirmer le nouveau mot de passe',
        value: '',
        updateable: updateMode,
      }"
      :inputType="'password'"
      :inputRules="[{ minLength: 5 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          passwordConfirmation = value;
        }
      "
    />

    <OnclickButton
      :onButtonClick="
        async () => {
          if (updateMode) {
            await updateUserPassword();
          }

          updateMode = !updateMode;
        }
      "
      :text="updateMode ? 'Enregistrer' : 'Modifier'"
      :disabled="
        (updateMode && updateButtonDisabled) ||
        (updateMode &&
          (!currentPassword || !newPassword || !passwordConfirmation))
      "
    />

    <OnclickButton
      v-if="updateMode"
      :onButtonClick="
        () => {
          updateMode = false;
          updateButtonDisabled = false;
          currentPassword = '';
          newPassword = '';
          passwordConfirmation = '';
        }
      "
      style="background-color: #f44336; margin-left: 10px"
      text="Annuler"
    />
  </BasicCard>
</template>
