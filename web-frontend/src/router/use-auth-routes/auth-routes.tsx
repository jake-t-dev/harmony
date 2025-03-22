import { useRoutes } from "react-router-dom";
import { useAuthRoutes } from "./use-auth-routes";

const AuthRoutes = () => {
    const routes = useAuthRoutes();
    const routeTree = useRoutes(routes);

    return routeTree;
}

export default AuthRoutes;