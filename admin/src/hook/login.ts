import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { LoginInfoState } from "../model/user";
import { getLoginInfo } from "../service/login";

export function useLogin() {
  const [loginInfo, setLoginInfo] = useRecoilState(LoginInfoState)

  useEffect(() => {
    getLoginInfo().then((res: any) => {
      setLoginInfo(res)
    }).catch(() => {})
  }, [])

  return { loginInfo }
}