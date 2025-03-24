import { Box, Divider } from "@mui/material";
import GroupIconBox from "../group-icon/group-icon-box";
import MenuRoundedIcon from "@mui/icons-material/MenuRounded";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";

type GroupListProps = { handleMenuClick?: () => void };

const GroupList = ({ handleMenuClick }: GroupListProps) => {
  return (
    <Box
      width={"80px"}
      bgcolor={"secondary.dark"}
      height={"calc(100vh - margin)"}
      margin={"20px"}
      sx={{
        borderRadius: "16px",
      }}
    >
      <GroupIconBox onClick={handleMenuClick}>
        <MenuRoundedIcon />
      </GroupIconBox>

      <Divider
        orientation="horizontal"
        variant="middle"
        sx={{
          marginY: "20px",
          backgroundColor: "secondary.light",
          borderRadius: 2,
        }}
      />

      <GroupIconBox
        overridingStyles={{
          bgcolor: "primary.main",
          transition: "border-radius 0.3s ease",
          borderRadius: "16px",
          "&:hover": {
            borderRadius: "60px",
            cursor: "pointer",
          },
        }}
      >
        <PeopleAltOutlinedIcon />
      </GroupIconBox>
    </Box>
  );
};

export default GroupList;
