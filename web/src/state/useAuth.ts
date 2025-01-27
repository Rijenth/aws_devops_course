import { reactive } from "vue";

type AuthState = {
  authenticated: boolean;
  token: string | null;
  username?: string | undefined;
  user_id?: string | undefined;
};

const state = reactive<AuthState>({
  authenticated: false,
  token: null,
  username: undefined,
  user_id: undefined,
});

const login = (
  token: string,
  username: string | undefined,
  userId: string | undefined,
) => {
  state.authenticated = true;
  state.token = token;
  state.username = username;
  state.user_id = userId;

  sessionStorage.setItem("token", state.token || "");
  sessionStorage.setItem("username", state.username || "");
  sessionStorage.setItem("user_id", state.user_id || "");
};

const logout = () => {
  sessionStorage.removeItem("token");
  sessionStorage.removeItem("username");
  sessionStorage.removeItem("user_id");

  state.authenticated = false;
  state.token = null;
  state.username = undefined;
  state.user_id = undefined;
};

export function useAuth() {
  const token = sessionStorage.getItem("token");
  const username = sessionStorage.getItem("username");
  const user_id = sessionStorage.getItem("user_id");

  if (!state.authenticated) {
    if (token && username && user_id) {
      login(token, username, user_id);
    }
  }

  return {
    state,
    login,
    logout,
  };
}
