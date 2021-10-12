import React from "react";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import {makeStyles, Typography} from "@material-ui/core";
import {CloseRounded, ZoomOutMapRounded} from "@material-ui/icons";
import ZoomOutRoundedIcon from '@material-ui/icons/ZoomOutRounded';

import {MasButton} from "./MasButton";
const useStyles = makeStyles((theme) => ({
  dialog: {
    width: (props: Iprops) => props.width,
    color: "red",
  },
  dialogTitle: {
    padding: "16px",
    borderBottom: "1px solid #ddd",
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },
  title: {
    fontSize: "1.6rem",
    fontFamily: "Helvetica-Bold, sans-serif",
  },
  headActions: {
    color: theme.palette.grey[600],
  },
  headIcons: {
    padding: "4px",
    cursor: "pointer",
  },
  body: {
    borderBottom: "1px solid #ddd",
    padding: 16,
  },
  footer: {
    // marginTop: theme.spacing(2),
    padding: "16px",
    justifyContent: "flex-start",
  },
}));
export interface Iprops {
  open: boolean;
  onClose: () => void;
  title: string;
  children?: any;
  onSave?: () => void;
  formType: string;
  width?: "xs" | "sm" | "md" | "lg" | "xl";
  onExit?: () => void;
  btonLabel?: string;
}

export function MasModal(props: Iprops) {
  const {
    open,
    title,
    onSave,
    formType,
    children,
    onClose,
    width,
    onExit,
    btonLabel,
  } = props;
  const classes = useStyles(props);
  const [full, setFull] = React.useState(false);
  const btnLabel = btonLabel
    ? btonLabel
    : formType === "new"
    ? "حفظ"
    : "تعديل"

  return (
    <Dialog
      open={open}
      maxWidth={width || "xs"}
      fullWidth={true}
      className={classes.dialog}
      onClose={onClose}
      fullScreen={full}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
      disableBackdropClick
      onExited={onExit}
    >
      <DialogTitle
        disableTypography
        className={classes.dialogTitle}
        id="alert-dialog-title"
      >
        <Typography className={classes.title} variant="h4">
          {title}
        </Typography>
        <div className={classes.headActions}>
          <span className={classes.headIcons}>
            {full?<ZoomOutRoundedIcon onClick={() => {
                setFull(!full);
              }}></ZoomOutRoundedIcon>
            :<ZoomOutMapRounded
              onClick={() => {
                setFull(!full);
              }}
            />}
          </span>
          <span className={classes.headIcons}>
            <CloseRounded onClick={onClose} />
          </span>
        </div>
      </DialogTitle>
      <DialogContent className={classes.body}>{children}</DialogContent>
      <DialogActions disableSpacing className={classes.footer}>
        {formType != "view" && formType != "" ? (
          <MasButton
            disableElevation={true}
            color="primary"
            type="submit"
            variant="contained"
            label={btnLabel}
            onClick={onSave}
          ></MasButton>
        ) : null}
        <MasButton
          disableFocusRipple={true}
          disableRipple={true}
          disableElevation={true}
          color="default"
          type="button"
          variant="outlined"
          label="إلغاء"
          onClick={onClose}
        />
      </DialogActions>
    </Dialog>
  );
}
