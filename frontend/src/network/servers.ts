const servers = ["http://localhost:8888/"];

export const getOneServer = () => {
  const id = ~~(Math.random() * servers.length);
  return servers[id] ?? servers[0];
};
