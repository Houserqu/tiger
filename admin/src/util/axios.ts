
import { toast } from 'amis';
import Axios from 'axios'
import history from '../history';

const axios = Axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL
})

axios.interceptors.response.use((response) => {
  const {
    data,
    config: {
      errMsg = true, // 错误提示信息（用于覆盖接口返回的提示信息）
      successMsg = false, // 成功提示信息（用于覆盖接口返回的提示信息）
      meta = false, // 是否返回完整响应体（用于需要自己判断成功失败的情况）
    },
  } = response;

  // 返回完整响应体
  if (meta) {
    return data;
  }

  // 业务逻辑错误
  if (data.errno !== 0) {
    if(data.errno === 402) {
      history.replace('/admin/login')
      return;
    }

    if (errMsg) {
      toast.error(typeof errMsg === 'string' ? errMsg : data.msg || '系统繁忙', {
        title: '操作失败'
      })
    }

    throw new Error(data.msg)
  }

  if (successMsg) {
    toast.success(typeof successMsg === 'string' ? successMsg : data.msg || '操作成功')
  }

  return data.data;
}, (err) => {
  // 网络错误
  toast.error(err.message || '系统繁忙', {
    title: `网络异常（${err?.response?.status}）`
  })

  return null;
})

export default axios;