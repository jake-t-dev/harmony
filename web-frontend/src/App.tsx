import { BrowserRouter } from "react-router-dom";
import ClientProvider from "./api/provider/client-provider";
import ErrorScreen from "./components/error-screen/error-screen";
import Routes from "./router/router";

export const App = () => {
  return (
    <ErrorScreen>
      <ClientProvider>
        <BrowserRouter>
          <Routes />
        </BrowserRouter>
      </ClientProvider>
    </ErrorScreen>
  );
};
