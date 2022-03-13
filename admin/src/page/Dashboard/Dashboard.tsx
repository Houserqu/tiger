import { Button } from "amis";
import renderAmis from "../../util/renderAmis";

export default function Dashboard({ children }: any) {
  return (
    <div>
      {renderAmis({
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