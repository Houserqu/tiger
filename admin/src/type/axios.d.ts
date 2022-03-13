import { AxiosRequestConfig } from 'axios';

declare module 'axios' {
  interface AxiosRequestConfig {
    meta?: boolean; // 是否返回完整 response
    successMsg?: boolean | string; // 请求成功提示内容：string = 提示，并作为提示文案，false = 不提示，true = 提示，但是用接口 msg 作为文案
    errMsg?: boolean | string; // 请求失败提示内容：string = 提示，并作为提示文案，false = 不提示，true = 提示，但是用接口 msg 作为文案
  }
}
