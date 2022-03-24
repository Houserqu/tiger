import { atom } from "recoil";

export const PagesState = atom<Page[]>({
  key: 'PagesState',
  default: []
})