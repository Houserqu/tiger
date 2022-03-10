import { render } from 'amis';
import fetcher from './fetcher';

/**
 * amis 渲染函数封装
 * @param schema 页面配置
 * @param props data
 * @param env 全局配置
 * @returns 
 */
export default function renderAmis(schema: any, props?: any, env?: any) {
  return render(schema, props, env || {
    fetcher
  })
}