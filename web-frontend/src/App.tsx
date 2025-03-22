import { BrowserRouter, Routes } from "react-router-dom";
import ClientProvider from "./api/provider/client-provider";
import ErrorScreen from "./components/error-screen/error-screen";

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
