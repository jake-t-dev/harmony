import { useMemo } from "react";
import { RouteObject } from "react-router-dom";
import Register from "../../pages/register/register";
import Login from "../../pages/login/login";
import Dashboard from "../../pages/dashboard/dashboard";

export const useAuthRoutes = () => {
  const authRoutes: RouteObject[] = useMemo(() => {
    return [
      {
        path: "/",
        element: <Dashboard />,
        children: [
          {
            path: "/login",
            element: <Login />,
          },
          {
            path: "/register",
            element: <Register />,
          },
        ],
      },
    ];
  }, []);
  return authRoutes;
};
