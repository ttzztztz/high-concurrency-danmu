import { IDanmuData } from "../types/type";
import { getOneServer } from "./servers";

export const sendDanmuToServer = async (item: IDanmuData) => {
  const perform_date = Date.now();
  const res = await fetch(`${getOneServer()}danmu/send`, {
    method: "POST",
    body: JSON.stringify(item),
    headers: new Headers({
      "Content-Type": "application/json",
    }),
  });

  const now_timestamp = Date.now();
  console.log(
    `[PERFORMANCE] send danmu ${now_timestamp - perform_date}ms`,
    item
  );
  return res;
};

export const getOnlineUser = async (rid: string) => {
  const res = await fetch(`${getOneServer()}danmu/online/${rid}`, {
    method: "GET",
  });

  return res;
};
