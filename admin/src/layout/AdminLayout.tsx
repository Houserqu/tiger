import { Avatar } from "amis";
import { Outlet } from "react-router-dom";
import renderAmis from "../util/renderAmis";

export default function AdminLayout({ children }: any) {
  return (
    <div>
      <div
        className="bg-white flex justify-between px-4 items-center"
        style={{ height: '60px', borderBottom: 'solid 1px rgb(231 231 231)'}}
      >
        <div>
          <i className='fab fa-html5 text-2xl' />
        </div>
        <Avatar icon='fa fa-user'/>
      </div>
      <div style={{ height: 'calc(100vh - 60px)', display: 'flex' }}>
        <div className='h-full bg-bule-100 bg-white'>
          {renderAmis({
            "type": "page",
            "body": {
              "type": "nav",
              "stacked": true,
              "className": "w-md",
              "links": [
                {
                  "label": "Nav 1",
                  "to": "/docs/index",
                  "icon": "fa fa-user",
                  "active": true
                },
                {
                  "label": "Nav 2",
                  "unfolded": true,
                  "children": [
                    {
                      "label": "Nav 2-1",
                    },
                    {
                      "label": "Nav 2-2",
                      "to": "/docs/api-2-2"
                    }
                  ]
                },
                {
                  "label": "Nav 3",
                  "to": "/docs/renderers"
                }
              ]
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