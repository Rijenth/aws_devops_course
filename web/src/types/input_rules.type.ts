export type InputRules =
  | "required"
  | "email"
  | "phone"
  | { minLength: number }
  | { maxLength: number };
