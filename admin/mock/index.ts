import { MockMethod } from 'vite-plugin-mock'
import { response } from './mockUtil'

const loginInfo = {
  userInfo: { id: 1, name: 'tom' },
  menu: [],
  permissions: []
}

const mock: MockMethod[] = [
  {
    url: '/api/user/loginInfo',
    method: 'get',
    response: response(loginInfo)
  },
  {
    url: '/api/user/login',
    method: 'post',
    response: response(loginInfo),
  },
  {
    url: '/api/config/page',
    method: 'get',
    response: response(loginInfo),
  },
]

export default mock