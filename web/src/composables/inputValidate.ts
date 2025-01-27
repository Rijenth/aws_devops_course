import { InputRules } from "../types/input_rules.type";

export default function useInputValidate(
  value: string,
  inputRules: InputRules[],
) {
  if (!inputRules) {
    throw new Error("Rules are required");
  }

  if (inputRules.length <= 0) {
    throw new Error("Rules cannot be empty");
  }

  for (const rule of inputRules) {
    if (rule === "required" && !value) {
      return "Ce champ est requis.";
    }

    if (rule === "email" && !isEmail(value as string)) {
      return "Ce champ doit être un email valide.";
    }

    if (rule === "phone" && !isValidPhoneNumber(value as string))
      return "Ce champ doit être un numéro de téléphone valide.";

    if (typeof rule === "object" && "minLength" in rule) {
      if (typeof value === "string" && value.length < rule.minLength)
        return `Ce champ doit contenir au moins ${rule.minLength} caractères.`;

      return;
    }

    if (typeof rule === "object" && "maxLength" in rule) {
      if (typeof value === "string" && value.length > rule.maxLength)
        return `Ce champ doit contenir au plus ${rule.maxLength} caractères.`;

      return;
    }
  }
}

function isEmail(email: string) {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
}

function isValidPhoneNumber(number: string) {
  if (!/^[0-9+-]*$/.test(number)) {
    return false;
  }

  return true;
}
