import {Grid2} from "@mui/material";
import ProfileIcon from "../profile-icon/profile-icon";

const Header = () => {
    return (
        <Grid2 container
            direction="row"
            justifyContent="flex-end"
            alignItems="center"
            padding={2}>
            <ProfileIcon />
        </Grid2>
    );
}

export default Header;