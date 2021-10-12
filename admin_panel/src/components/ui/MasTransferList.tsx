import React, { useEffect,  useState } from 'react';
import List, { ListProps } from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import * as icon from '@material-ui/icons';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import InputLabel from '@material-ui/core/InputLabel';
// import  {ReactSortable } from 'react-sortablejs';


export interface MasTransferListProps extends ListProps {
    items: Array<any>,
    name?: string,
    displayMember: string,
    valueMember: string,
    selectedlist: any[],
    titleSelected: string,
    titleUnSelected: string,
    onChange: (data: any) => void
}



export function MasTransferList(ownProp: MasTransferListProps) {
    const { items, onChange, name, titleSelected, titleUnSelected, className, displayMember, valueMember } = ownProp;
    let { selectedlist } = ownProp
    const MasListItemss: Array<any> = items || []
    const [, setValulist] = useState({})
    let [sortedList,setSortedList]=useState(selectedlist)

    useEffect(()=>{
        let temp:Array<any>=[]
        selectedlist.map(el=>{             
         let item=MasListItemss.filter(itm=>itm[valueMember]===el)
         if(item.length>0)
            {
                item[0]['id']=el
                temp.push(item[0])
         }
     })
     setSortedList(temp)
    },[selectedlist])


    const handleSelectItem = (value: string | React.FormEvent<HTMLUListElement>, child?: undefined) => {
        selectedlist?.push(value);
        let newItem = MasListItemss?.filter(itm => itm[valueMember] === value)
        if (newItem.length > 0) {
            newItem[0]['id'] = value
        }
        setValulist(value + "False")
        if (onChange != null) {
            onChange(selectedlist)
        }
    };

    const handleDeSelectItem = (value: string | React.FormEvent<HTMLUListElement>, child?: undefined) => {
        console.log(value,'deselectd')
        let index: number = selectedlist?.indexOf(value) || -1;
            selectedlist?.splice(index, 1);
            setValulist(value + "True")        
            if (onChange != null) {
            onChange(selectedlist)
            }
    };
    const handleDragEnd=()=>{
        let afterSorted=sortedList.map(item=>{return item[valueMember]} )
        console.log(afterSorted)
        selectedlist=afterSorted
        // if(onDragEnd){
        //     onDragEnd(afterSorted)
        // }
       }
    return (
        <table className={className}>
            <tbody>
                <tr>
                    <td style={{ textAlign: 'right', verticalAlign: 'top', }}>
                        <InputLabel htmlFor={"selectedfor" + name}>{titleSelected}</InputLabel>
                        <List>
                            {
                                MasListItemss ? selectedlist ? MasListItemss.filter(cls => !selectedlist?.includes(cls[valueMember]))
                                    .map((item, index) => {
                                        return (
                                            <ListItem onDoubleClick={e => handleSelectItem(item[valueMember])} id={index.toString()} key={index} button>
                                                <ListItemAvatar>
                                                    <Avatar >
                                                        <icon.MoreHoriz />
                                                    </Avatar>
                                                </ListItemAvatar>
                                                <ListItemText style={{ textAlign: 'right' }} primary={item[displayMember]} />
                                            </ListItem>
                                        )
                                    }) : <ListItem key={0} /> : <ListItem key={0} />
                            }
                        </List></td>
                    <td style={{ textAlign: 'right', verticalAlign: 'top', }}>
                        <InputLabel htmlFor={"unselectedfor" + name}>{titleUnSelected}</InputLabel>
                        <List id="items">
                            {
                                
                                    MasListItemss?selectedlist?MasListItemss.filter(cls => selectedlist.includes(cls[valueMember]))
                                 .map((item,index) => {
                                        return (
                                            <ListItem className="item"
                                                onDoubleClick={e => handleDeSelectItem(item[valueMember])}
                                                key={index} button>
                                                <ListItemAvatar>
                                                    <Avatar >
                                                        <icon.MoreHoriz />
                                                    </Avatar>
                                                </ListItemAvatar>
                                                <ListItemText style={{ textAlign: 'right' }} primary={item[displayMember]} />
                                            </ListItem>

                                        )
                                    })
                                    : <ListItem key={0} /> : <ListItem key={0} />
                              }
                              
                             
                        </List>
                    </td>
                </tr></tbody></table>
    );
}