import { useRoutes } from "react-router-dom";
import AdminLayout from "./layout/AdminLayout";
import AmisEditor from "./page/AmisEditor/AmisEditor";
import Dashboard from "./page/Dashboard/Dashboard";
import Index from "./page/Index/Index";
import NotFound from "./page/NotFound";

export const routes = [
  {
    path: "/",
    element: <Index />,
  },
  {
    path: "/editor",
    element: <AmisEditor />,
  },
  {
    path: "/admin",
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
