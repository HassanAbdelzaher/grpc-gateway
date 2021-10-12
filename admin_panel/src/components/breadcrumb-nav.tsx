import React from 'react';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import { makeStyles, useTheme, } from '@material-ui/core/styles'
import Link from '@material-ui/core/Link';
import { grey } from '@material-ui/core/colors'

function handleClick(event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) {
    event.preventDefault();
    console.info('You clicked a breadcrumb.');
}
const useStyles = makeStyles((theme) => ({
    root: {
        textAlign: "right",
        borderBottom: `1px solid ${grey[300]}`,
        backgroundColor: "#fff",
        width: "100%",
        padding: theme.spacing(1, 2),
    }
}));
export default function NavBreadcrumbs() {
    const classes = useStyles();
    const theme = useTheme();
    return (
        <Breadcrumbs className={classes.root} separator="/" aria-label="breadcrumb">
            <Link color="inherit" href="/" onClick={handleClick}>
                الصفحة الرئيسية
             </Link>
            <Link color="inherit" href="/getting-started/installation/" onClick={handleClick}>
                ادارة النظام
            </Link>
            <Link
                style={{ color: theme.palette.common.black }}
                href="/"
                onClick={handleClick}
                aria-current="page"
            >  الوظائف
            </Link>
        </Breadcrumbs>
    );
}