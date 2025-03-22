import { useMemo } from "react"
import { RouteObject } from "react-router-dom";

export const useAuthRoutes = () => {
    const authRoutes: RouteObject[] = useMemo(() => {
        return [
            {
                path: "/",
                //element: <Home />,
            },
        ];
    }
    , []);
    return authRoutes;
}