import { getOneServer } from "./servers";
import { IDanmuData } from "../types/type";

export const connectToWebsocket = (
  uid: number,
  rid: number,
  danmuHandler: (data: IDanmuData) => any
) => {
  const url = `${getOneServer().replace(/^http/, "ws")}ws/${uid}/${rid}`;
  const socket = new WebSocket(url);

  socket.onopen = () => {
    console.log("socket connected");
  };

  socket.onmessage = (message: any) => {
    const { data } = message;
    danmuHandler(JSON.parse(data));
  };

  socket.onerror = (err: any) => {
    console.error("socket err", err);
  };

  return socket;
};
