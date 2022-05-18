import { render, toast } from "amis";
import _ from "lodash";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { useRecoilState } from "recoil";
import { PagesState } from "../model/page";
import { getPageByPath } from "../service/config";
import fetcher from "../util/fetcher";

export function useRenderAmis() {

  /**
   * amis 渲染函数封装
   * @param schema 页面配置
   * @param props data
   * @param env 全局配置
   * @returns
   */
  const renderAmis = (schema: any, props?: any, env?: any) =>
    render(
      schema,
      props,
      env || {
        fetcher,
      }
    );

  return [renderAmis];
}

export function usePageAmisConfig(): [boolean, Page | undefined] {
  const [loading, setLoading] = useState(false);
  const [pages, setPages] = useRecoilState(PagesState);
  const location = useLocation();

  // 根据路径查找页面配置
  let page = _.find(pages, { path: location.pathname });

  useEffect(() => {
    let page = _.find(pages, { path: location.pathname });

    if (!page) {
      setLoading(true);
      getPageByPath(location.pathname)
        .then((res) => {
          try {
            res.config = JSON.parse(res.config || "{}");
            res.extend = JSON.parse(res.extend || "{}");
            setPages([...pages, res]);
          } catch (error) {
            console.log(error);
            throw new Error("解析页面配置失败");
          }
        })
        .catch((err) => {
          toast.error(err.message, {
            title: "页面异常",
          });
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [location.pathname]);

  return [loading, page];
}
