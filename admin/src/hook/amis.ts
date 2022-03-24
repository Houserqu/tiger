import { render } from 'amis';
import _ from 'lodash';
import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { PagesState } from '../model/page';
import { getPageByPath } from '../service/page';
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
      }
    })
  }

  return [renderAmis]
}

export function usePageAmisConfig(): [boolean, Page | undefined] {
  const [loading, setLoading] = useState(false);
  const [pages, setPages] = useRecoilState(PagesState)
  const location = useLocation()

  // 根据路径查找页面配置
  let page = _.find(pages, { path: location.pathname })

  useEffect(() => {  
    let page = _.find(pages, { path: location.pathname })

    if(!page) {
      setLoading(true)
      getPageByPath(location.pathname).then((res) => {
        setPages([...pages, res])
      })
      .catch(() => {})
      .finally(() => { setLoading(false) })
    }
  }, [location.pathname])

  return [loading, page]
}