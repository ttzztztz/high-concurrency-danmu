import { IDanmuData } from "../types/type";
import { getOneServer } from "./servers";

export const sendDanmuToServer = async (item: IDanmuData) => {
  const res = await fetch(`${getOneServer()}danmu/send`, {
    method: "POST",
    body: JSON.stringify(item),
    headers: new Headers({
      "Content-Type": "application/json",
    }),
  });

  return res
};
