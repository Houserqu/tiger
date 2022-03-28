import { useRoutes } from "react-router-dom";
import AdminLayout from "./layout/AdminLayout";
import AmisPage from "./page/AmisPage/AmisPage";
import Dashboard from "./page/Dashboard/Dashboard";
import Login from "./page/Login/Login";
import NotFound from "./page/NotFound";

export const routes = [
  {
    path: "/login",
    element: <Login />,
  },
  {
    element: <AdminLayout />,
    children: [
      { index: true, element: <Dashboard /> },
      { path: "*", element: <AmisPage /> } // 未匹配到的路径，则代表通过 amis 加载，
    ],
  },
  { path: "*", element: <NotFound /> }
];

function MyRouter() {
  return useRoutes(routes);
}

export default MyRouter;
