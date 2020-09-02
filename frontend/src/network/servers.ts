require('dotenv').config()
const SERVER_0 = process.env.BUILD === "prod" ? "http://10.108.21.47:31888/" : "http://localhost:8888/";

const servers = [SERVER_0];
export const getOneServer = () => {
  return servers[0];
};
