import React from "react";
import {withStyles, Theme} from "@material-ui/core/styles";
import Tooltip from "@material-ui/core/Tooltip";
import Zoom from "@material-ui/core/Zoom";

interface _props {
  title: string;
  children?: any;
}

const CustomeTooltip = withStyles((theme: Theme) => ({
  tooltip: {
    backgroundColor: "rgb(0 0 0 / 70%)",
    color: "#fff",
    boxShadow: theme.shadows[1],
    fontSize: 12,
    padding: "6px",
  },
  arrow: {
    color: "rgb(0 0 0 / 70%)",
  },
}))(Tooltip);

export default function MasTooltip(props: _props) {
  const {title, children} = props;
  return (
    <CustomeTooltip title={title} TransitionComponent={Zoom} arrow>
      {children}
    </CustomeTooltip>
  );
}
