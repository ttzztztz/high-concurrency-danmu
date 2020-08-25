import { getOneServer } from "./servers";
import { IDanmuData } from "../types/type";

export const connectToWebsocket = (
  uid: number,
  rid: number,
  danmuHandler: (data: IDanmuData) => any,
  errHandler: (data: any) => any
) => {
  const url = `${getOneServer().replace(/^http/, "ws")}ws/${uid}/${rid}`;
  const socket = new WebSocket(url);

  socket.onopen = () => {
    console.log("socket connected");
  };

  socket.onmessage = (message: any) => {
    const { data } = message;

    // todo: 粘包问题
    try {
      danmuHandler(JSON.parse(data));
    } catch (e) {
      console.error(e, data)
    }
  };

  socket.onerror = (err: any) => {
    console.error("socket err", err);
    errHandler(err)
  };

  return socket;
};
