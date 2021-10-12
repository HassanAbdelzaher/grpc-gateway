import React, { useState} from "react";
import Header from "./header";
import {makeStyles} from "@material-ui/core/styles";
import Routes from "../routes/routs";
const width = "260px";

const useStyles = makeStyles((theme) => ({
  root: {
    width: "100%",
    height: "100%",
  },
  appContainer: {
    width: `calc(100% - ${theme.spacing(9)}px)`,
    height: "100%",
    //padding: theme.spacing(1),
    // overflow: "hidden",
    position: "relative",
    marginTop: theme.spacing(9),
    marginLeft: theme.spacing(9),
    transition: theme.transitions.create("all", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  appContainerShift: {
    width: `calc(100% - ${width})`,
    marginLeft: `${width}`,
    transition: theme.transitions.create("all", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appHeader: {},
  title: {
    fontSize: "1.6rem",
    fontFamily: "Helvetica-Bold",
  },
}));

export default function Layout() {
  const [open] = useState(false);
  const classes = useStyles();

  const handleClicked = () => {
    //setOpen((op) => !op);
  };
  
   
  return (
    <div className={classes.root}>
      <Header
        drawerWidth={width}
        titleStyle={classes.title}
        title=" test"
        handleDrawerOpen={handleClicked}
        open={open}
      />
        <div>
        <Routes />
        </div>
    </div>
  );
}
