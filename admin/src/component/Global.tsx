import {ToastComponent, AlertComponent} from 'amis';
import { useLogin } from '../hook/login';

export default function Global({children}: any) {
  useLogin()
  return (
    <div>
      <ToastComponent key="toast" />
      <AlertComponent key="alert" />
      {children}
    </div>
  )
}