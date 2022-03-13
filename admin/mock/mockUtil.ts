export function response(data: any, errno = 0) {
  return {
    errno,
    data,
    msg: '成功'
  }
}