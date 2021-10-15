import React,{useEffect} from 'react'
import {MasList,MasButton} from '../components/ui'
import { IconButton, makeStyles } from '@material-ui/core'
import {Link} from 'react-router-dom'
import axios from 'axios'
import TableApi from '../components/table'
import * as Icons  from '@material-ui/icons'
let configurl='./configpane.json'
const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: "80%",
        padding:10,
        borderRadius:5,
        margin:'auto'
    },
    list: {
        width: '60%',
        background:'white',
        //direction:'rtl'
    }
}))

export const Main=()=>{
    const classes = useStyles()

const [state,setState]=React.useState<any>(null)
  useEffect(()=>{
     axios.get(configurl).then(res=>{
         setState([res.data])
     })      
  },[])
    return (
        <div className={classes.root}> 
           <div style={{display:'inline-block'}}> <h2> Servies Names </h2> </div>
           <div style={{display:'inline-block'}}>
           <IconButton  style={{padding: "4px", paddingBottom: "0"}} >
           <Link to='/addConfig'> 
           <Icons.AddCircleOutlineOutlined/>
           </Link>
			</IconButton>
           </div>
            <div>
                <TableApi data={state} columns={['host','http_port']} />
            </div>
            {/* <MasList listStyle={classes.list} name='services' itemValue='service_name' primaryItemText='service_name' items={state?.services}/> */}
            <br/>
        </div> 
    )
}