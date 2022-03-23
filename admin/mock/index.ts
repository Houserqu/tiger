import { MockMethod } from 'vite-plugin-mock'
import { response } from './mockUtil'

const loginInfo = {
  userInfo: { id: 1, name: 'tom' },
  menu: [
    { id: 'dashboard', parent_id: '', label: '工作台', to: '/admin', icon: 'fa fa-user', target: '', permission: '', },
    { id: 'product', parent_id: '', label: '物品', to: '', icon: 'fa fa-user', target: '', permission: '' },
    { id: 'productList', parent_id: 'product', label: '物品列表', to: '/admin/product/list', icon: 'fa fa-user', target: '', permission: '' },
    { id: 'user', parent_id: '', label: '用户', to: '', icon: 'fa fa-user', target: '', permission: '' },
    { id: 'userList', parent_id: 'user', label: '用户列表', to: '/admin/user/list', icon: 'fa fa-user', target: '_blank', permission: '' },
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