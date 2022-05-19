import axios from "../util/axios";

export function getPageByPath(path: string): Promise<Page> {
  return axios.get('/api/page/detail', { params: { path } })
}