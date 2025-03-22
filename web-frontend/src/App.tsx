import { BrowserRouter } from "react-router-dom";
import ClientProvider from "./api/provider/client-provider";
import ErrorScreen from "./components/error-screen/error-screen";
import Routes from "./router/router";
import { ThemeProvider } from "@emotion/react";
import theme from "./theme";
import { CssBaseline } from "@mui/material";

export const App = () => {
  return (
    <ErrorScreen>
      <ThemeProvider theme={theme}>
      <CssBaseline />
        <ClientProvider>
          <BrowserRouter>
            <Routes />
          </BrowserRouter>
        </ClientProvider>
      </ThemeProvider>
    </ErrorScreen>
  );
};
