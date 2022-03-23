import { useRoutes } from "react-router-dom";
import AdminLayout from "./layout/AdminLayout";
import AmisEditor from "./page/AmisEditor/AmisEditor";
import Dashboard from "./page/Dashboard/Dashboard";
import Login from "./page/Login/Login";
import NotFound from "./page/NotFound";

export const routes = [
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/editor",
    element: <AmisEditor />,
  },
  {
    element: <AdminLayout />,
    children: [
      { index: true, element: <Dashboard /> },
      { path: "*", element: <NotFound /> }
    ],
  },
  { path: "*", element: <NotFound /> }
];

function MyRouter() {
  return useRoutes(routes);
}

export default MyRouter;
