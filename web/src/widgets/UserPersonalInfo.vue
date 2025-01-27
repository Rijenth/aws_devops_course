<script setup lang="ts">
import { defineProps, onMounted, ref } from "vue";
import { ConvertTimestampToHumanReadable } from "../composables/ConvertTimestampToHumanReadable";
import OnclickButton from "../components/OnclickButton.vue";
import EditableParagraph from "../components/EditableParagraph.vue";
import UserModel from "../models/user.model";
import DatePicker from "../components/DatePicker.vue";
import { Timestamp } from "../generated/google/protobuf/timestamp";
import UserApi from "../api/user.api";
import { useAuth } from "../state/useAuth";
import router from "../router/router";
import BasicCard from "../components/BasicCard.vue";
import { useToast } from "../state/useToast";

const props = defineProps<{
  updateMode?: boolean;
  username?: string;
}>();

const toast = useToast();
const userApi = new UserApi();
const auth = useAuth();

const user = ref<UserModel | null>(null);
const allowUpdate = ref<boolean>(false);
const updateButtonDisabled = ref<boolean>(false);

const patchUser = async () => {
  if (!user.value) {
    toast.notify(
      "Impossible de mettre à jour les informations de l'utilisateur, aucun utilisateur fourni",
      "error",
    );
    fetchAndFillUser();
    return;
  }

  let currentUsername = auth.state.username;

  const response = await userApi.patch(user.value.updateUserRequestFormat());

  if (typeof response === "string") {
    toast.notify(response, "error");
    fetchAndFillUser();
    return;
  }

  if (!response.user) {
    toast.notify(
      "Impossible de récupérer les informations de l'utilisateur depuis le serveur",
      "error",
    );
    fetchAndFillUser();
    return;
  }

  user.value = new UserModel(response.user);

  if (currentUsername !== user.value.username) {
    auth.logout();
    router.push({ name: "Accueil" });
    return;
  }
};

const fetchAndFillUser = async () => {
  let username = props.username ?? auth.state.username;

  if (!username) {
    toast.notify(
      "Impossible de récupérer les informations de l'utilisateur, aucun nom d'utilisateur fourni",
      "error",
    );
    return;
  }

  const response = await userApi.getByUsername(username);

  if (typeof response === "string") {
    toast.notify(response, "error");
    return;
  }

  if (!response.user) {
    toast.notify(
      "Impossible de récupérer les informations de l'utilisateur depuis le serveur",
      "error",
    );
    return;
  }

  user.value = new UserModel(response.user);
};

onMounted(async () => {
  fetchAndFillUser();
});
</script>

<template>
  <BasicCard :title="'Informations personnelles'">
    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Nom d\'utilisateur',
        value: user.username,
        updateable: allowUpdate,
      }"
      :inputRules="['required', { minLength: 5 }, { maxLength: 20 }]"
      :information-bubble="true"
      :information-message="`Pour des raisons de sécurité, toute modification de ce champ entraînera une déconnexion de votre session`"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.username = value;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Nom',
        value: user.lastName,
        updateable: allowUpdate,
      }"
      :inputRules="['required', { minLength: 2 }, { maxLength: 20 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.lastName = value;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Prénom',
        value: user.firstName,
        updateable: allowUpdate,
      }"
      :inputRules="['required', { minLength: 2 }, { maxLength: 20 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.firstName = value;
          }
        }
      "
    />

    <DatePicker
      v-if="user"
      title="Né(e) le"
      :timestamp="user.dateOfBirth"
      :updateable="allowUpdate"
      @updated="
        (timestamp: Timestamp) => {
          if (user) {
            user.dateOfBirth = timestamp;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Adresse mail',
        value: user.email,
        updateable: allowUpdate,
      }"
      :inputType="'email'"
      :inputRules="['required', 'email']"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.email = value;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Adresse',
        value: user.address,
        updateable: allowUpdate,
      }"
      :inputRules="['required']"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.address = value;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Numéro de téléphone',
        value: user.phoneNumber,
        updateable: allowUpdate,
      }"
      :inputType="'tel'"
      :inputRules="['required', 'phone']"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.phoneNumber = value;
          }
        }
      "
    />

    <EditableParagraph
      v-if="user"
      :params="{
        title: 'Biographie',
        value: user.bio,
        updateable: allowUpdate,
      }"
      :inputRules="['required', { minLength: 10 }, { maxLength: 20 }]"
      @rulesNotMet="updateButtonDisabled = $event"
      @updated="
        (value: string) => {
          if (user) {
            user.bio = value;
          }
        }
      "
    />

    <p>
      <strong>Inscrit depuis le </strong>
      {{ ConvertTimestampToHumanReadable(user?.createdAt) }}
    </p>

    <OnclickButton
      v-if="updateMode"
      :onButtonClick="
        async () => {
          if (allowUpdate) await patchUser();
          allowUpdate = !allowUpdate;
        }
      "
      :text="allowUpdate ? 'Enregistrer' : 'Modifier'"
      :disabled="updateButtonDisabled"
    />

    <OnclickButton
      v-if="allowUpdate"
      :onButtonClick="
        () => {
          fetchAndFillUser();
          allowUpdate = false;
        }
      "
      style="background-color: #f44336; margin-left: 10px"
      text="Annuler"
    />
  </BasicCard>
</template>

<style>
p {
  margin: 10px 0;
  line-height: 1.6;
}
</style>
