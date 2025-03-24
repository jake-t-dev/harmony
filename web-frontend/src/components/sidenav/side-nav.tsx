import { Divider, Drawer, Grid2 } from "@mui/material";
import { useState } from "react";
import GroupList from "./group-list/group-list";

const SideNav = () => {
  const [isOpen, setIsOpen] = useState(false);

  const handleToggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <Grid2
        container
        direction="row"
        sx={{ width: "300px", backgroundColor: "secondary.main" }}
      >
        <GroupList />
        <Drawer anchor="left" open={isOpen} onClose={handleToggle}>
          <p>SideNav</p>
        </Drawer>
      </Grid2>
      <Divider
        orientation="vertical"
        sx={{
          height: "100vh",
          backgroundColor: "text.secondary",
        }}
      />
    </>
  );
};

export default SideNav;
