import { render } from 'amis';
import { useNavigate } from 'react-router-dom';
import fetcher from '../util/fetcher';

export function useRenderAmis() {
  const navigate = useNavigate()
  
  /**
   * amis 渲染函数封装
   * @param schema 页面配置
   * @param props data
   * @param env 全局配置
   * @returns 
   */
  function renderAmis(schema: any, props?: any, env?: any) {
    return render(schema, props, env || {
      fetcher,
      jumpTo(to, action) {
        navigate(to)
        console.log(to)
        console.log(action)
      }
    })
  }

  return [renderAmis]
}
