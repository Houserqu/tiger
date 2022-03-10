import { Outlet } from "react-router-dom";

export default function AdminLayout({children}: any) {
  return (
    <div>
      <Outlet />
    </div>
  )
}