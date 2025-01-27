import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router/router";
import Vue3Toastify from "vue3-toastify";

createApp(App)
  .use(Vue3Toastify, {
    position: "top-right",
    theme: "light",
  })
  .use(router)
  .mount("#app");
