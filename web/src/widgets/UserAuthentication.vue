<script setup lang="ts">
import { ref } from "vue";
import OnclickButton from "../components/OnclickButton.vue";
import { useAuth } from "../state/useAuth";
import LabelledInput from "../components/LabelledInput.vue";
import AuthApi from "../api/auth.api";
import { jwtDecode } from "jwt-decode";
import { useToast } from "../state/useToast";

const auth = useAuth();
const toast = useToast();
const authApi = new AuthApi();
const username = ref<string>("");
const password = ref<string>("");

const handleLogin = async () => {
  if (!username.value || !password.value) {
    toast.notify(
      "Veuillez entrer un nom d'utilisateur et un mot de passe",
      "error",
    );
    return;
  }

  const response = await authApi.login(username.value, password.value);

  if (typeof response !== "string") {
    const decodedToken = response.token
      ? (jwtDecode(response.token) as { name?: string; sub?: string })
      : null;

    auth.login(response.token, decodedToken?.name, decodedToken?.sub);

    return;
  }

  toast.notify(response, "error");
};
</script>

<template>
  <div class="login-box">
    <h2>Connexion</h2>

    <LabelledInput
      label="Nom d'utilisateur"
      id="username"
      type="text"
      v-model="username"
      placeholder="Entrez votre nom d'utilisateur"
    />

    <LabelledInput
      label="Mot de passe"
      id="password"
      type="password"
      v-model="password"
      placeholder="Entrez votre mot de passe"
    />

    <OnclickButton :onButtonClick="handleLogin" text="Se connecter" />
  </div>
</template>

<style scoped>
.login-box {
  display: flex;
  flex-direction: column;
  align-items: left;
  min-width: 300px;
  padding: 20px;
  border: 1px solid #7289da;
  border-radius: 8px;
}

h2 {
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  font-size: 18px;
  color: #666;
}

input {
  width: 100%;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ddd;
  font-size: 14px;
  box-sizing: border-box;
}
</style>
