import {ToastComponent, AlertComponent} from 'amis';
import { useLogin } from '../hook/login';
import Loading from './Loading';

export default function Global({children}: any) {
  const { loginUserInfo } = useLogin()
  console.log(loginUserInfo);
  return (
    <div className="min-h-full">
      <ToastComponent key="toast" />
      <AlertComponent key="alert" />
      {loginUserInfo ? children : (
        <div style={{height: '100vh'}}>
          <Loading text='登录中' />
        </div>
      )}
    </div>
  )
}