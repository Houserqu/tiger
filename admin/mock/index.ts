import { MockMethod } from 'vite-plugin-mock'
import { response } from './mockUtil'

const loginInfo = {
  userInfo: { id: 1, name: 'tom' },
  menu: [
    { id: 1, parent_id: 0, label: '工作台', to: '/', icon: 'fa fa-user', target: '', permission: '', },
    { id: 2, parent_id: 0, label: '物品', to: '', icon: 'fa fa-user', target: '', permission: '' },
    { id: 3, parent_id: 2, label: '物品列表', to: '/product/list', icon: 'fa fa-user', target: '', permission: '' },
    { id: 4, parent_id: 0, label: '用户', to: '', icon: 'fa fa-user', target: '', permission: '' },
    { id: 5, parent_id: 4, label: '用户列表', to: '/user/list', icon: 'fa fa-user', target: '_blank', permission: '' },
  ],
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