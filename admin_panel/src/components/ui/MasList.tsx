import React, { useEffect } from 'react';
import List, { ListProps } from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import InputLabel from '@material-ui/core/InputLabel';
import { Link } from 'react-router-dom'
export interface MasListProps extends ListProps {
    items: any[],
    name: string,
    itemValue?: string,
    labelStyle?: any,
    label?: string,
    avatarIcon?: any
    icons?: Array<any> //array of icons
    primaryItemText: string
    SecondaryItemText?: string,
    listStyle?: string,
    itemStyle?: string,
    iconStyle?: string,
    textStyle?: string,
    linkStyle?: any,
    onClick?:any
}

export function MasList(ownProp: MasListProps) {
    const {onClick, listStyle, itemStyle, iconStyle, textStyle,linkStyle, label, items, icons, itemValue, avatarIcon, primaryItemText, SecondaryItemText, labelStyle, name } = ownProp
    const MasListItemss: any[] = items
    const [ItemList, setListItem] = React.useState(1)
    const handleListClick = (item: any,index) => {
        setListItem(index)
        sessionStorage.setItem(name,index)
        if(onClick){
            onClick(item,index)
        }
    }
    useEffect(()=>{
        let item:any=sessionStorage.getItem(name)||0
        setListItem(parseInt(item))
    },[name])

    const renderList = () => {
        return (
            <>
                <InputLabel className={labelStyle} htmlFor={"htmlfor" + name}>{label}</InputLabel>
                <List className={listStyle}>
                    {MasListItemss ? MasListItemss.map((item, index) => {
                        return (
                            item.to ? <Link to={item.to} className={linkStyle} key={index}>
                                <ListItem className={itemStyle} key={index} button
                                onClick={() => handleListClick(item,index)}
                                selected={ItemList === index ? true : false}
                            >
                                {
                                    avatarIcon ? <ListItemAvatar>
                                        <Avatar>
                                            {avatarIcon}
                                        </Avatar>
                                    </ListItemAvatar> : null
                                }
                                {icons ?
                                    <ListItemIcon className={iconStyle}>
                                        {icons[index]}
                                    </ListItemIcon> : null}
                                <ListItemText
                                    className={textStyle}
                                    primary={item[primaryItemText]}
                                    secondary={SecondaryItemText ? { SecondaryItemText } : null}
                                />
                            </ListItem>
                            </Link> : (
                                    <ListItem className={itemStyle} key={index} button
                                        onClick={() => handleListClick(item,index)}
                                        selected={ItemList === index ? true : false}
                                    >
                                        {
                                            avatarIcon ? <ListItemAvatar>
                                                <Avatar>
                                                    {avatarIcon}
                                                </Avatar>
                                            </ListItemAvatar> : null
                                        }
                                        {icons ?
                                            <ListItemIcon className={iconStyle}>
                                                {icons[index]}
                                            </ListItemIcon> : null}
                                        <ListItemText
                                            className={textStyle}
                                            primary={item[primaryItemText]}
                                            secondary={SecondaryItemText ? { SecondaryItemText } : null}
                                        />
                                    </ListItem>
                                )
                        )
                    }) : <ListItem key={0} />
                    }
                </List>
            </>
        )
    };
    return (
        <>
            {renderList()}
        </>
    );
}