const SERVER_0 = process.env.NODE_ENV === "production" ? "http://101.133.130.112:31888/" : "http://localhost:8888/";

const servers = [SERVER_0];
export const getOneServer = () => {
  return servers[0];
};
