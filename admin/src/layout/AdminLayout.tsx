import { Avatar } from "amis";
import { Outlet } from "react-router-dom";
import _ from 'lodash';
import { useLoginInfo } from "../hook/login";
import { useMemo } from "react";
import { useRenderAmis } from "../hook/amis";

// 生成菜单树
function genMenuTree(root: number, data: { id: number, parent_id: number}[]): any[] {
  const list: any[] = _.filter(data, { parent_id: root })

  return list.map((v) => {
    const children = genMenuTree(v.id, data);

    if(children.length > 0) {
      return { ...v, children }
    }

    return v
  })
}

export default function AdminLayout({ children }: any) {
  const { loginMenu } = useLoginInfo()
  const [render] = useRenderAmis()
  const menuData = useMemo(() => genMenuTree(0, loginMenu), [loginMenu])

  return (
    <div>
      <div
        className="bg-white flex justify-between px-4 items-center"
        style={{ height: '60px', borderBottom: 'solid 1px rgb(231 231 231)' }}
      >
        <div>
          <i className='fab fa-html5 text-2xl' />
        </div>
        <Avatar icon='fa fa-user' />
      </div>
      <div style={{ height: 'calc(100vh - 60px)', display: 'flex' }}>
        <div className='h-full bg-bule-100 bg-white'>
          {render({
            "type": "page",
            "body": {
              "type": "nav",
              "stacked": true,
              "className": "w-md",
              "links": menuData
            }
          })}
        </div>
        <div className="w-full bg-gray-100">
          <Outlet />
        </div>
      </div>
    </div>
  )
}