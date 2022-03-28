import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { LoginMenuState, LoginPermissionsState, LoginUserInfoState } from "../model/user";
import { getLoginInfo } from "../service/login";

export function useLogin() {
  const [loginUserInfo, setLoginUserInfo] = useRecoilState(LoginUserInfoState)
  const [loginMenu, setLoginMenu] = useRecoilState(LoginMenuState)
  const [loginPermissions, setLoginPermissions] = useRecoilState(LoginPermissionsState)

  useEffect(() => {
    getLoginInfo().then((res: any) => {
      setLoginUserInfo(res.userInfo)
      setLoginMenu(res.menus)
      setLoginPermissions(res.permissions)
    }).catch(() => {})
  }, [])

  return { loginUserInfo, loginMenu, loginPermissions }
}

export function useLoginInfo() {
  const [loginUserInfo] = useRecoilState(LoginUserInfoState)
  const [loginMenu] = useRecoilState(LoginMenuState)
  const [loginPermissions] = useRecoilState(LoginPermissionsState)

  return { loginUserInfo, loginMenu, loginPermissions }
}