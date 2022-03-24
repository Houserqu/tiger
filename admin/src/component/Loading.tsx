import { Spinner } from "amis";

type Props = {
  text?: string
}
export default function Loading({ text = '加载中...' }: Props) {
  return (
    <div className="flex justify-center items-center h-full">
      <Spinner size='large' />
      <div className="ml-2">{text}</div>
    </div>
  )
}