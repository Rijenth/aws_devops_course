<script setup lang="ts">
import { computed, ref } from "vue";
import UserPersonalInfo from "../widgets/UserPersonalInfo.vue";
import UserPasswordEdit from "../widgets/UserPasswordEdit.vue";

const navBarItems = ["Informations personnelles", "Changer mon mot de passe"];
const selectedNavIndex = ref<number>(0);
const selectedNavItem = computed(() => navBarItems[selectedNavIndex.value]);
</script>

<template>
  <div class="container">
    <div class="main">
      <div class="user-details">
        <h2 class="nav-title">Navigation</h2>
        <nav>
          <ul>
            <li
              v-for="(navItem, index) in navBarItems"
              :key="index"
              @click="selectedNavIndex = index"
              :style="{ color: selectedNavIndex === index ? '#7289da' : '' }"
            >
              {{ navItem }}
            </li>
          </ul>
        </nav>
      </div>

      <UserPersonalInfo
        v-if="selectedNavItem === 'Informations personnelles'"
        :updateMode="true"
      />

      <UserPasswordEdit
        v-else-if="selectedNavItem === 'Changer mon mot de passe'"
      />
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  color: #333;
}

.main {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.user-details {
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  min-width: 250px;
  max-height: 160px;
}

.user-details nav ul {
  display: flex;
  flex-direction: column;
  text-align: left;
  gap: 15px;
}

.nav-title {
  font-size: 1.5em;
  margin-bottom: 15px;
  color: #7289da;
  text-decoration: underline;
}

ul {
  list-style-type: none;
}

li {
  cursor: pointer;
}

li:hover {
  text-decoration: underline;
  color: #7289da;
}

@media screen and (min-width: 768px) {
  .main {
    flex-direction: row;
    gap: 40px;
  }

  .user-personal-info {
    width: 100%;
  }
}
</style>
