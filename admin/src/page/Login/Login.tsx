import { useRenderAmis } from "../../hook/amis"

export default function Index() {
  const [render] = useRenderAmis()

  return (
    <div style={{height: '100vh'}} className="flex h-full justify-center items-center">
      {render({
        "type": "page",
        "className": 'w-1/4 h-auto',
        "body": {
          "type": "form",
          "title": "登录",
          redirect: '/',
          "api": {
            "method": "post",
            "url": "/api/login/loginByPhone",
            "data": {
              "&": "$$"
            }
          },
          "body": [
            {
              "name": "phone",
              "type": "input-text",
              "label": "手机号"
            },
            {
              "name": "password",
              "type": "input-password",
              "label": "密码"
            }
          ]
        }
      })}
    </div>
  )
}