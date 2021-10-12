import React from 'react';
import { Button, makeStyles } from '@material-ui/core';
import clsx from 'clsx';

export interface Btnprops {
    label: string
    color?: "primary" | "secondary" | "default";
    type: 'button' | 'submit'
    variant: "contained" | "outlined" | "text"
    size?: 'small' | "large" | 'medium'
    styleBtn?: any,
    disabled?: boolean,
    disableElevation?: boolean,
    disableFocusRipple?: boolean,
    disableRipple?: boolean
    fullWidth?: boolean,
    startIcon?: any,
    onClick?: () => void
}
const useStyles = makeStyles((theme) => ({
    root: {
        display: 'inline-block',
        marginRight: "12px",
        fontSize: '1.2rem',
        padding: theme.spacing(0.5, 4)
    }
}))


export function MasButton(_props: Btnprops) {
    const { startIcon, color, onClick, label, type, variant, size, styleBtn, fullWidth, disableElevation, disableFocusRipple, disableRipple, disabled, ...restProps } = _props
    const classes = useStyles();
    return (
        <Button
            className={clsx(classes.root, styleBtn)}
            fullWidth={fullWidth}
            disabled={disabled}
            disableElevation={disableElevation}
            disableFocusRipple={disableFocusRipple}
            disableRipple={disableRipple}
            size={size}
            type={type}
            color={color}
            variant={variant}
            onClick={onClick}
            endIcon={startIcon}
            {...restProps}>
            {label}
        </Button>
    );
}