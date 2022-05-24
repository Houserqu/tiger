import { render, toast } from "amis";
import _ from "lodash";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { useRecoilState } from "recoil";
import history from "../history";
import { PagesState } from "../model/page";
import { getPageByPath } from "../service/config";
import fetcher from "../util/fetcher";

/**
 * amis 路由配置
 * https://github.com/baidu/amis/blob/master/examples/app/index.jsx
 * @returns 
 */

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
        updateLocation: (location: string, replace: boolean) => {
          console.log(history);
          location = normalizeLink(location);
          if (location === "goBack") {
            return history.back();
          } else if (
            !/^https?\:\/\//.test(location) &&
            location === history.location.pathname + history.location.search
          ) {
            // 目标地址和当前地址一样，不处理，免得重复刷新
            return;
          } else if (/^https?\:\/\//.test(location) || !history) {
            return (window.location.href = location);
          }

          history[replace ? "replace" : "push"](location);
        },
        jumpTo: (to, action) => {
          if (to === "goBack") {
            return history.back();
          }

          to = normalizeLink(to);

          if (isCurrentUrl(to)) {
            return;
          }

          if (action && action.actionType === "url") {
            action.blank === false
              ? (window.location.href = to)
              : window.open(to, "_blank");
            return;
          } else if (action && action.blank) {
            window.open(to, "_blank");
            return;
          }

          if (/^https?:\/\//.test(to)) {
            window.location.href = to;
          } else if (
            (!/^https?\:\/\//.test(to) &&
              to === history.location.pathname + history.location.search)
          ) {
            // do nothing
          } else {
            history.push(to);
          }
        },
        isCurrentUrl: isCurrentUrl,
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

function normalizeLink(to: string, location = history.location) {
  to = to || "";

  if (to && to[0] === "#") {
    to = location.pathname + location.search + to;
  } else if (to && to[0] === "?") {
    to = location.pathname + to;
  }

  const idx = to.indexOf("?");
  const idx2 = to.indexOf("#");
  let pathname = ~idx
    ? to.substring(0, idx)
    : ~idx2
    ? to.substring(0, idx2)
    : to;
  let search = ~idx ? to.substring(idx, ~idx2 ? idx2 : undefined) : "";
  let hash = ~idx2 ? to.substring(idx2) : location.hash;

  if (!pathname) {
    pathname = location.pathname;
  } else if (pathname[0] != "/" && !/^https?\:\/\//.test(pathname)) {
    let relativeBase = location.pathname;
    const paths = relativeBase.split("/");
    paths.pop();
    let m;
    while ((m = /^\.\.?\//.exec(pathname))) {
      if (m[0] === "../") {
        paths.pop();
      }
      pathname = pathname.substring(m[0].length);
    }
    pathname = paths.concat(pathname).join("/");
  }

  return pathname + search + hash;
}

function isCurrentUrl(to: string, ctx?: any) {
  if (!to) {
    return false;
  }
  const pathname = history.location.pathname;
  const link = normalizeLink(to, {
    ...location,
    pathname,
    hash: "",
    state: null,
    key: ''
  });

  // if (!~link.indexOf("http") && ~link.indexOf(":")) {
    // return match(link, {
    //   decode: decodeURIComponent,
    //   strict: typeof ctx?.strict !== "undefined" ? ctx.strict : true,
    // })(pathname);
  // }

  return decodeURI(pathname) === link;
}
