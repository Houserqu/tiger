import renderAmis from "../../util/renderAmis";

export default function Index() {
  return (
    <div style={{height: '100vh'}} className="flex h-full justify-center items-center">
      {renderAmis({
        "type": "page",
        "className": 'w-1/4 h-auto',
        "body": {
          "type": "form",
          "title": "登录",
          redirect: '/admin',
          "api": {
            "method": "post",
            "url": "/pub/login",
            "data": {
              "&": "$$"
            }
          },
          "body": [
            {
              "name": "name",
              "type": "input-text",
              "label": "账号"
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