import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  shape: {
    borderRadius: 8,
  },
  palette: {
    primary: {
      main: "#7E22CE",
      light: "#E9D5FF",
    },
    secondary: {
      main: "#18181B",
      light: "#D4D4D8",
      dark: "#09090B",
    },
    background: {
      default: "#262626",
      paper: "#262626",
    },
    text: {
      primary: "#fff",
      secondary: "#52525B",
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
