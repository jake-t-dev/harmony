import { Box, Button, Divider } from "@mui/material";
import AppDrawer from "./app-drawer";
import SpaTwoToneIcon from "@mui/icons-material/SpaTwoTone";
import PeopleIcon from "@mui/icons-material/People";

const SideNav = () => {
    return (
        <AppDrawer>
            <Box
                sx={{
                    display: "flex",
                    justifyContent: "center",
                    margin: "15px 0",
                }}
            >
                <SpaTwoToneIcon sx={{ fontSize: 40, color: "primary.main" }} />
            </Box>
            <Divider variant="middle" />
            <Button
                variant="contained"
                sx={{ maxWidth: "80%", justifyContent: "center", margin: "10px 10px" }}
            >
                <PeopleIcon />
            </Button>
        </AppDrawer>
    );
};

export default SideNav;
