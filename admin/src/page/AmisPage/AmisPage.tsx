import Loading from "../../component/Loading"
import { usePageAmisConfig, useRenderAmis } from "../../hook/amis"
import NotFound from "../NotFound"

export default function AmisPage({children}: any) {
  const [loading, page] = usePageAmisConfig()
  const [render] = useRenderAmis()

  if(loading) {
    return <Loading />
  }

  if(!page) {
    return <NotFound />
  }

  return (
    <div>{render(page?.config)}</div>
  )
}