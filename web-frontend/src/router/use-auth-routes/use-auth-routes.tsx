import { useMemo } from "react";
import { RouteObject } from "react-router-dom";
import DefaultLayout from "../../components/default-layout/default-layout";
import GroupChat from "../../pages/group-chat/group-chat";

export const useAuthRoutes = () => {
  const authRoutes: RouteObject[] = useMemo(() => {
    return [
      {
        path: "/",
        element: <DefaultLayout />,
        children: [
          {
            path: "/chat/:id",
            element: <GroupChat />,
          },
        ],
        
      },
    ];
  }, []);
  return authRoutes;
};
