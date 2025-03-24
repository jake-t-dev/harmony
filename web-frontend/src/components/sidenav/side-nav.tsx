import { Box, Divider, Grid2 } from "@mui/material";
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
        sx={{
          transition: "width 0.3s ease",
          width: isOpen ? "300px" : "120px",
          backgroundColor: "secondary.main",
        }}
      >
        <GroupList handleMenuClick={handleToggle} />
        <Box>
          <Box
            width={isOpen ? "160px" : "0px"}
            height={"60px"}
            borderRadius={"16px"}
            sx={{
              position: "absolute",
              bottom: 25,
              bgcolor: "secondary.dark",
              transition: "width 0.3s ease",
            }}
          ></Box>
        </Box>
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
