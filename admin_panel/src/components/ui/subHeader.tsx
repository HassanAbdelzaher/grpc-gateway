import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Box, Typography } from '@material-ui/core';
import clsx from 'clsx';
import Liquid from '../../assets/images/svg/Liquid-Cheese.svg'

const useStyles = makeStyles((theme) => ({
    root: {
        textAlign: "center",
        padding: "24px",
        paddingBottom: '4px',
        borderBottom: '1px solid #d7dde1',
        position: 'relative',
        '&:after': {
            content: "''",
            background: `url(${Liquid})`,
            top: 0,
            left: 0,
            position: 'absolute',
            zIndex: - 1,
            backgroundRepeat: "no-repeat",
            backgroundSize: "cover",
            width: "100%",
            height: "100%"
        },
        '& h1': {
            display: "inline-block",
            fontFamily: "Helvetica-Bold",
            fontSize: "2.2rem",
            position: "relative",
            padding: "4px",
            color:"#fff"
        },
        '& p': {
            fontSize: "1.1rem",
            padding: "4px",
            color:'#1b3c66'
        }
    }
}));
interface subHeader {
    title?: string
    subtitle?: string
    headerStyle?: any
}
export default function SubHeader(props: subHeader) {
    const { title, headerStyle, subtitle } = props
    const classes = useStyles(props);
    return (
        <Box className={clsx(classes.root, { headerStyle })} >
            <Typography variant="h3" component="h1" >
                {title}
            </Typography>
            {subtitle ? <Typography variant="subtitle1" component="p">{subtitle}</Typography> : null}
        </Box>
    );
}