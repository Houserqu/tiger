import axios from "../util/axios";

export function getLoginInfo() {
  return axios.get('/api/user/loginInfo')
}