import { useMemo } from "react";
import { RouteObject } from "react-router-dom";
import Register from "../../pages/register/register";
import Login from "../../pages/login/login";
import DefaultLayout from "../../components/default-layout/default-layout";

export const useAuthRoutes = () => {
  const authRoutes: RouteObject[] = useMemo(() => {
    return [
      {
        path: "/",
        element: <DefaultLayout />,
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
