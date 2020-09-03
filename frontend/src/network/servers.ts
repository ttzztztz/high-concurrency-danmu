const SERVER_0 = process.env.NODE_ENV === "production" ? "http://10.108.21.47:31888/" : "http://101.133.209.70:8888/";

const servers = [SERVER_0];
export const getOneServer = () => {
  return servers[0];
};
