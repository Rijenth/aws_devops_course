import { createRouter, createWebHistory } from "vue-router";
import { useAuth } from "../state/useAuth";

export const routes = [
  {
    path: "/",
    name: "Accueil",
    component: () => import("../pages/HomePage.vue"),
    meta: {
      addToNavbar: true,
      requiresAuth: false,
    },
  },
  {
    path: "/users-list",
    name: "Mes utilisateurs",
    component: () => import("../pages/UsersListPage.vue"),
    meta: {
      addToNavbar: true,
      requiresAuth: true,
    },
  },
  {
    path: "/my-account",
    name: "Mon compte",
    component: () => import("../pages/UserAccountPage.vue"),
    meta: {
      addToNavbar: true,
      requiresAuth: true,
    },
  },
  {
    path: "/logout",
    name: "Déconnexion",
    component: () => import("../pages/LogoutPage.vue"),
    meta: {
      addToNavbar: true,
      requiresAuth: true,
    },
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("../pages/NotFoundPage.vue"),
    meta: {
      addToNavbar: false,
      requiresAuth: false,
    },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _, next) => {
  const auth = useAuth();
  const requiresAuth = to.meta.requiresAuth;

  if (!requiresAuth) {
    next();
    return;
  }

  if (!auth.state.authenticated) {
    next({ name: "Accueil" });
    return;
  }

  next();
});

export default router;
