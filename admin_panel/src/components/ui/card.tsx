import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
//import Cardliq from '../../assets/images/svg/cardLiq.svg';
import { Link } from 'react-router-dom'
import { Card, CardActions, CardContent, Typography, CardHeader } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    card: {
        display: 'flex',
        position: 'relative',
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        textAlign: 'center',
        margin: theme.spacing(1),
        padding: theme.spacing(1),
        borderRadius: '8px',
        zIndex: 1,
        cursor: 'pointer',
        [theme.breakpoints.down('sm')]: {
            margin: theme.spacing(1.2),
        },
    },
    overlay: {
        position: 'absolute',
        top: '100px',
        left: '30px',
        width: "120%",
        height: '100%',
        //backgroundImage: `url(${Cardliq})`,
        backgroundRepeat: 'no-repeat',
        backgroundSize: '80%',
        opacity: '50%',
        zIndex: -1
    },
    link: {
        textDecoration: 'none',
        color: "inherit",
        position: "absolute",
        top: 0,
        left: 0,
        width: "100%",
        height: "100%"
    }
}));
interface cardprops {
    title: any,
    icon?: React.ReactElement //icon 
    discription?: string,
    elevation?: number,
    styleCard?: string,
    styleHeader?: string,
    styleTitle?: string,
    styleDiscription?: string,
    styleContent?: string,
    to?: any,
    onClick? :any

}
export function MasCard(props: cardprops) {
    const {onClick, to, styleContent, elevation, title, icon, discription, styleCard, styleHeader, styleTitle, styleDiscription } = props
    const classes = useStyles();

    return (
        <Card onClick={onClick}  elevation={elevation} className={clsx(classes.card, [styleCard])} raised>
            {to ? <Link className={classes.link} to={to}></Link> : null}
            <CardHeader
                className={styleHeader}
                avatar={icon ? icon : null} />
            <CardContent className={styleContent}>
                <Typography variant="h3" className={styleTitle}>
                    {title}
                </Typography>
            </CardContent>
            <CardActions className={styleDiscription} >
                {discription}
            </CardActions>
            <div className={classes.overlay}></div>
        </Card>
    );
}
