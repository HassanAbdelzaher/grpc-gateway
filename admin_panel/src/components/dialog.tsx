import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import {Warning} from '@material-ui/icons'
//import { useTranslation } from 'react-i18next';
export enum DeleteDialogResult{
    CANCELLED=0,
    OK=1
}
export default function DeleteDialog(props:{open:boolean,onClose?:(resul:DeleteDialogResult)=>void}) {
  const {open} = props;
  //const { t } = useTranslation();

  const handleOk = () => {
    if(props.onClose){
        props.onClose(DeleteDialogResult.OK)
    }
  };

  const handleClose = () => {
    if(props.onClose){
        props.onClose(DeleteDialogResult.CANCELLED)
    }
  };

  return (
    <div>
      <Dialog open={open} disableBackdropClick aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">
          Warning           
          <Warning style={{color:'orange',float:'left'}} fontSize="large"/>
        </DialogTitle>
        <DialogContent>
          <DialogContentText>
            <hr></hr>
            <div>Are you sure you want to delete the selected rows?</div>
          </DialogContentText>
        </DialogContent>
        <hr></hr>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleOk} color="primary">
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}