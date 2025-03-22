import { Drawer } from "@mui/material";

type AppDrawerProps = {
  children: React.ReactNode;
};

const AppDrawer = ({ children }: AppDrawerProps) => {
  return (
    <Drawer
      anchor="left"
      variant="permanent"
      sx={{ width: "100%", maxWidth: 240,}}
    >
      {children}
    </Drawer>
  );
};

export default AppDrawer;