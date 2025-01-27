import { Timestamp } from "../generated/google/protobuf/timestamp";

export const ConvertTimestampToHumanReadable = (
  timestamp: Timestamp | undefined,
) => {
  if (timestamp === undefined) {
    return "";
  }

  const date = new Date(Number(timestamp.seconds) * 1000);

  const day = String(date.getUTCDate()).padStart(2, "0");
  const month = String(date.getUTCMonth() + 1).padStart(2, "0");
  const year = date.getUTCFullYear();

  return `${day}-${month}-${year}`;
};
