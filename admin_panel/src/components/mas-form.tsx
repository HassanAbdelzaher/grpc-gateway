import React from 'react'
import { MasButton, MasTextField } from './ui'
import { makeStyles, Typography } from '@material-ui/core'
import { CloseRounded, ZoomOutMapRounded } from '@material-ui/icons';
const useStyles = makeStyles((theme) => ({
    form: {
        margin: "100px",
        backgroundColor: "#fff",
        padding: 20,
        borderRadius: 8,
    },
    head: {
        display: 'flex',
        justifyContent: "space-between",
        alignItems: "center",
        padding: "4px 8px",
        borderBottom: "1px solid #eee",
        // fontFamily: 'helvetica-bold'
    },
    title: {
        fontSize: "1.7rem"
    },
    headActions: {
        color: theme.palette.grey[600]
    },
    headIcons: {
        padding: "4px",
        cursor: "pointer"
    },
    body: {
        borderBottom: "1px solid #eee",
        paddingTop: 16,
        paddingBottom: 4
    },
    footer: {
        marginTop: theme.spacing(2)
    }
}))

export default function MasForm(props:any) {
    const {title,btnLabel,children,onSave,onCancel}=props
    const classes = useStyles()
    return (
        <div className={classes.form}>
            <div className={classes.head}>
                <Typography className={classes.title} variant="h4">{title}</Typography>
                <div className={classes.headActions}>
                    <span className={classes.headIcons}>
                        <CloseRounded />
                    </span>
                    <span className={classes.headIcons}>
                        <ZoomOutMapRounded />
                    </span>
                </div>
            </div>
            <div className={classes.body}>
               {children}
            </div>
            <div className={classes.footer}>
                <MasButton disableElevation={true} color="primary" type="submit" variant="contained" label={btnLabel} onClick={onSave}/>
                <MasButton disableFocusRipple={true} disableRipple={true} disableElevation={true} color="default" type="button" variant="outlined" label="إلغاء" onClick={onCancel} />
            </div>

        </div>
    )
}
