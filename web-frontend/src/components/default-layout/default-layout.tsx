import { styled } from "@mui/material"
import { Outlet } from "react-router-dom"
import SideNav from "../sidenav/sidenav"
import Header from "../header/header"

const DefaultLayout = () => {
    return (
        <Container>
            <SideNav />
            <Page>
                <Header />
                <Content>
                    <Outlet />
                </Content>
            </Page>
        </Container>
    )
}

export default DefaultLayout

const Container = styled("div")({
    width: "100%",
    height: "100vh",
    overflow: "auto",
    display: "flex",
})

const Page = styled("div")({
    flex: "auto",
    display: "flex",
    flexDirection: "column",
    height: "100%",
    overflow: "auto",
})

const Content = styled("div")({
    flex: "auto",
    padding: "20px",
    height: "100%",
    overflow: "auto",
})