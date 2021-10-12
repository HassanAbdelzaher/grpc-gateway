import React,{useEffect} from 'react'
import {MasList,MasButton} from '../components/ui'
import { makeStyles } from '@material-ui/core'
import {Link} from 'react-router-dom'
import axios from 'axios'
let configurl='./configpane.json'
const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: "50%",
        margin: " 30px auto",
        boxShadow: "0, 0, 30px rgba(0,0,0,0.4)" ,
        background:'rgb(173 181 226 / 7%)',
        padding:10,
        border:'1px solid',
        borderRadius:5
        

    },
    list: {
        width: '60%',
        //direction:'rtl'
    }
}))

export const Main=()=>{
    const classes = useStyles()

const [state,setState]=React.useState<any>(null)
  useEffect(()=>{
     axios.get(configurl).then(res=>{
         setState(res.data)
     })      
  },[])
    return (
        <div className={classes.root}> 
            <h2 style={{fontSize:30}}> Servies Names</h2>

            <MasList listStyle={classes.list} name='services' itemValue='service_name' primaryItemText='service_name' items={state?.services}/>
            <br/>
            <Link to='/addConfig'> <MasButton label='add New Config' variant='contained' type="button" color='primary' /></Link>
        </div> 
    )
}