import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { Box } from "@mui/material";

const ProfieIcon = () => {
  return (
    <Box
      sx={{
        width: 50,
        height: 50,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        borderRadius: "50%",
        transition: "background-color 0.3s ease, transform 0.2s ease",
        "&:hover": {
          backgroundColor: "primary.main",
          boxShadow: 1,
          transform: "scale(1.1)",
        },
        cursor: "pointer",
      }}
    >
      <AccountCircleIcon
        sx={{
          fontSize: 40,
          color: "primary.main",
          "&:hover": {
            color: "white",
            },
        }}
      />
    </Box>
  );
};

export default ProfieIcon;
