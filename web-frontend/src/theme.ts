import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  shape: {
    borderRadius: 8,
  },
  palette: {
    primary: {
      main: "#7e22ce",
      light: "#d8b4fe",
    },
    secondary: {
      main: "#3f3f46",
    },
    background: {
      default: "#27272a",
      paper: "#27272a",
    },
    text: {
      primary: "#fff",
      secondary: "#a0a0a0",
    },
  },
  typography: {
    allVariants: {
      color: "#fff",
    },
    fontFamily: '"Tektur", "Roboto", "Arial", sans-serif',
  },
});

export default theme;
