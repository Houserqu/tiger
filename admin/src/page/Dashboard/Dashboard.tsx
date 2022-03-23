import { useRenderAmis } from "../../hook/amis";

export default function Dashboard({ children }: any) {
  const [render] = useRenderAmis()

  return (
    <div>
      {render({
        "type": "page",
        "body": {
          "type": "panel",
          "title": "首页",
          "body": "首页"
        }
      })}
    </div>
  )
}