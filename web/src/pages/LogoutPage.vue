// src/pages/LogoutPage.vue
<script setup lang="ts">
import { useAuth } from "../state/useAuth";
import { useRouter } from "vue-router";
import { onMounted } from "vue";
import AuthApi from "../api/auth.api";

const auth = useAuth();
const authApi = new AuthApi();
const router = useRouter();

onMounted(async () => {
  await authApi.logout();
  auth.logout();
  router.push({ name: "Accueil" });
});
</script>

<template>
  <p>Déconnexion en cours...</p>
</template>
