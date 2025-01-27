import { toast, ToastType } from "vue3-toastify";
import "vue3-toastify/dist/index.css";

export function useToast() {
  const notify = (message: string, type: ToastType = "success") => {
    if (type === "info") toast.info(message);
    if (type === "success") toast.success(message);
    if (type === "error") toast.error(message);
    if (type === "warning") toast.warning(message);
    if (type === "loading") toast.loading(message);
    if (type === "default") toast.success(message);
  };

  return { notify };
}
