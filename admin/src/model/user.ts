import { atom } from "recoil";

export const LoginUserInfoState = atom({
  key: 'LoginUserInfoState',
  default: null
})

export const LoginMenuState = atom({
  key: 'LoginMenuState',
  default: []
})

export const LoginPermissionsState = atom({
  key: 'LoginPermissionsState',
  default: []
})