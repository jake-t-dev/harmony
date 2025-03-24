import { Box, SxProps, Theme } from "@mui/material";

type GroupIconBoxProps = {
  children: React.ReactNode;
  overridingStyles?: SxProps<Theme>;
  onClick?: () => void;
};

const GroupIconBox = ({ children, overridingStyles, onClick }: GroupIconBoxProps) => {
  return (
    <Box
      width={"60px"}
      height={"60px"}
      borderRadius={"60px"}
      bgcolor={"secondary.main"}
      margin={"10px"}
      display="flex"
      justifyContent="center"
      alignItems="center"
      sx={{
        transition: 'border-radius 0.3s ease',
        "&:hover": {
          bgcolor: "primary.main",
          borderRadius: "16px",
          cursor: "pointer",
        },
        ...overridingStyles,
      }}
      onClick={onClick}
    >
      {children}
    </Box>
  );
};

export default GroupIconBox;
