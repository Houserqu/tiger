import { useRoutes } from "react-router-dom";
import AdminLayout from "./layout/AdminLayout";
import Dashboard from "./page/Dashboard/Dashboard";
import Home from "./page/Dashboard/Dashboard";
import NotFound from "./page/NotFound";

export const routes = [
  {
    path: "/",
    element: <Home />,
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
