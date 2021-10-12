import React, {} from "react";
import {makeStyles} from "@material-ui/core/styles";
import {AppBar, Typography, Toolbar, IconButton} from "@material-ui/core";
// import {Translator} from '../lang/lang'

interface headerProps {
  title?: string;
  drawerWidth: string;
  titleStyle: any;
  open: boolean;
  handleDrawerOpen: any;
}
// const x = (props: headerProps) => props.drawerWidth;

const useStyles = makeStyles((theme) => ({
  appBar: {
    padding: "4px !important",
    color: theme.palette.text.primary,
    boxShadow:
      "0 0.4rem 2.1875rem rgba(4,9,20,0.03), 0 0.7rem 1.40625rem rgba(4,9,20,0.03), 0 0.15rem 0.53125rem rgba(4,9,20,0.05), 0 0.125rem 0.1875rem rgba(4,9,20,0.03)",
    zIndex: theme.zIndex.drawer + 1,
    transition: "all 0.25s ease-out",
    background: "#fff",
  },
  appBarShift: {
    // marginRight: "260px",
    width: `calc(100% - 260px)`, // need to convert props.drawerWidth
    transition: "all 0.25s ease",
  },
  grow: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: 36,
  },
  hide: {
    display: "none",
  },
  sectionDesktop: {
    display: "none",
    [theme.breakpoints.up("md")]: {
      display: "flex",
    },
  },
  sectionMobile: {
    display: "flex",
    [theme.breakpoints.up("md")]: {
      display: "none",
    },
  },
  icons: {
    fontSize: "28px",
  },
  toolbar: {
    [theme.breakpoints.down("md")]: {
      minHeight: "64px",
    },
  },
}));

export default function Header(props: headerProps) {
    return (
    <>
      <AppBar position="static">
  <Toolbar>
    {/* <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
      <MenuIcon />
    </IconButton> */}
    <Typography variant="h6" >
      admin_panel
    </Typography>
  </Toolbar>
</AppBar>
     
    </>
  );
}
