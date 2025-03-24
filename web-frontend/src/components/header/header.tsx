import { Box, Divider, Typography } from "@mui/material";

const Header = () => {
  return (
    <>
      <Box height={"100px"}>
        <Typography variant="h5">Header</Typography>
      </Box>
      <Divider
        orientation="horizontal"
        sx={{
          width: "100%",
          backgroundColor: "text.secondary",
        }}
      />
    </>
  );
};

export default Header;
