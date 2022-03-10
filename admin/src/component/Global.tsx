import {ToastComponent, AlertComponent} from 'amis';

export default function Global({children}: any) {
  return (
    <div>
      <ToastComponent key="toast" />
      <AlertComponent key="alert" />
      {children}
    </div>
  )
}