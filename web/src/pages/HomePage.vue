<script setup lang="ts">
import { useAuth } from "../state/useAuth";
import UserAuthentication from "../widgets/UserAuthentication.vue";

const auth = useAuth();
</script>

<template>
  <div>
    <p class="info-message">
      {{
        auth.state.authenticated
          ? "Connecté en tant que " + auth.state.username
          : "Déconnecté"
      }}
    </p>

    <br />

    <UserAuthentication v-if="!auth.state.authenticated" />
  </div>
</template>

<style scoped>
.info-message {
  background-color: #f0f0f0;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  color: #666;
  font-size: 16px;
  text-align: center;
  max-width: 80%;
  margin: 0 auto;
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.1);
}
</style>
