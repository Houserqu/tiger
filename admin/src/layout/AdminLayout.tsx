import { Avatar } from "amis";
import { Outlet, useNavigate } from "react-router-dom";
import _ from "lodash";
import { useLoginInfo } from "../hook/login";
import { useMemo } from "react";
import { Menu as AntdMenu } from "antd";

// 生成菜单树
function genMenuTree(
  root: number,
  data: Menu[]
): any[] {
  const list: any[] = _.filter(data, { parent_id: root });

  return list.map((v) => {
    const children = genMenuTree(v.id, data);
    const menu: any = {
      label: v.label,
      key: v.id
    }

    if (children.length > 0) {
      menu.children = children
    }

    return menu;
  });
}

export default function AdminLayout({ children }: any) {
  const { loginMenu } = useLoginInfo();
  const navigate = useNavigate()
  const menuData = useMemo(() => genMenuTree(0, loginMenu), [loginMenu]);

  const changeMenu = ({ key }: { key: string }) => {
    const menu = _.find(loginMenu, { id: parseInt(key)})
    if(menu?.to) {
      navigate(menu.to)
    }
  }

  return (
    <div>
      <div
        className="bg-white flex justify-between px-4 items-center"
        style={{ height: "60px", borderBottom: "solid 1px rgb(231 231 231)" }}
      >
        <div>
          <i className="fab fa-html5 text-2xl" />
        </div>
        <Avatar icon="fa fa-user" />
      </div>
      <div style={{ height: "calc(100vh - 60px)", display: "flex" }}>
        <AntdMenu
          mode="inline"
          theme="dark"
          items={menuData}
          style={{ width: 300, height: "100%" }}
          onClick={changeMenu}
        />
        <div className="w-full bg-gray-100">
          <Outlet />
        </div>
      </div>
    </div>
  );
}
