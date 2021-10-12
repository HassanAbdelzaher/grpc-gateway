import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
// import {Link} from 'react-router-dom'

export interface props{
  onConfirmClick:any,
  open:boolean,
  handleClose:any,
  content:any,
  title:any,
  btnTitle?:string
}
export default function AlertDialog(props:props) {
    const {open,handleClose,onConfirmClick,content,title,btnTitle}=props

  return (
    <div>      
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{title}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description" style={{fontSize:25}}>
           {content}
          </DialogContentText>
        </DialogContent>
        {/* {actions? */}
        <DialogActions>
          <Button onClick={handleClose} color="primary" style={{fontSize:25}}>
          الغاء
          </Button>
          <Button onClick={onConfirmClick } color="primary" autoFocus style={{fontSize:25}}>
             حذف
          </Button>
        </DialogActions>
      {/* //   :
      //   <DialogActions>
      //   <Link style={{ textDecoration: 'none' }} to={to} ><Button onClick={onAddClick} color="primary" style={{fontSize:25}}>
      //   <Translator id={"اضافة"} values={{ count: 0 }} />
      //   </Button></Link>
      //   <Button onClick={onConfirmClick } color="primary" autoFocus style={{fontSize:25}}>
      //   <Translator id={"حذف"} values={{ count: 0 }} />
      //   </Button>
      // </DialogActions>} */}
      </Dialog>
    </div>
  );
}