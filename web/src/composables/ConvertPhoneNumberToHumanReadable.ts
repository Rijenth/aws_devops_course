export const ConvertPhoneNumberToHumanReadable = (
  phoneNumber: string | undefined,
) => {
  //TODO: refactor this to handle country prefix
  if (!phoneNumber) {
    return "";
  }

  if (phoneNumber.length !== 10) {
    return "";
  }

  return phoneNumber.replace(/(\d{2})(?=\d)/g, "$1.");
};
