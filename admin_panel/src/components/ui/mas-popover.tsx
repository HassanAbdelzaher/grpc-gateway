import React from 'react';
import Popover from '@material-ui/core/Popover';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        popover: {
            pointerEvents: 'none',
        },
        paper: {
            padding: theme.spacing(1),
        },
    }),
);
interface _props {
    open: boolean
    anchorEl: HTMLElement | null
    id?: string
    onClose?: () => void
    children: any
    elevation?: number | 0
}

export default function MasPopover(props: _props) {
    const { open, anchorEl, id, onClose, children, elevation } = props
    const classes = useStyles();

    return (
        <div>
            <Popover
                id={id}
                className={classes.popover}
                classes={{
                    paper: classes.paper,
                }}
                open={open}
                anchorEl={anchorEl}
                anchorOrigin={{
                    vertical: 'center',
                    horizontal: 'left',
                }}
                transformOrigin={{
                    vertical: 'center',
                    horizontal: 'right',
                }}
                onClose={onClose}
                elevation={elevation}
                disableRestoreFocus
            >
                {children}
            </Popover>
        </div>
    );
}
