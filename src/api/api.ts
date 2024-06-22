import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:1323/api/v1",
});

export default API;
