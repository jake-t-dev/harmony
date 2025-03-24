import { Box, SxProps, Theme } from "@mui/material";

type GroupIconBoxProps = {
  children: React.ReactNode;
  overridingStyles?: SxProps<Theme>;
};

const GroupIconBox = ({ children, overridingStyles }: GroupIconBoxProps) => {
  return (
    <Box
      width={"60px"}
      height={"60px"}
      borderRadius={"60px"}
      bgcolor={"primary.main"}
      margin={"10px"}
      display="flex"
      justifyContent="center"
      alignItems="center"
      sx={{
        transition: 'border-radius 0.3s ease',
        "&:hover": {
          borderRadius: "16px",
        },
        ...overridingStyles,
      }}
    >
      {children}
    </Box>
  );
};

export default GroupIconBox;
