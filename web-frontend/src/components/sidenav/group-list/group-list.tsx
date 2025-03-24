import { Box, Divider } from "@mui/material";
import GroupIconBox from "../group-icon/group-icon-box";
import MenuRoundedIcon from "@mui/icons-material/MenuRounded";
import PeopleAltOutlinedIcon from "@mui/icons-material/PeopleAltOutlined";

const GroupList = () => {
  return (
    <Box
      width={"80px"}
      bgcolor={"secondary.dark"}
      height={"calc(100vh - margin)"}
      margin={"20px"}
      sx={{
        borderRadius: 2,
      }}
    >
      <GroupIconBox
        overridingStyles={{
          backgroundColor: "secondary.main",
        }}>
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
          borderRadius: "16px",
          "&:hover": {
            borderRadius: "60px",
          },
        }}
      >
        <PeopleAltOutlinedIcon />
      </GroupIconBox>
    </Box>
  );
};

export default GroupList;
