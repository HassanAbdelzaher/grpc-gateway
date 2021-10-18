import React, { useState } from 'react'
import { makeStyles,Grid,TextField } from '@material-ui/core'
import { MasTextField,MasButton,MasCheckBox,MasModal } from '../components/ui'
import {AddService} from './add-service'
import {Link} from 'react-router-dom'
import { useForm } from "react-hook-form";

const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: "80%",
        margin: "auto",
        height:'100%',
        textAlign:'center',
    },
    styleFormControl: {
        width: '100%'
    },
    btn:{
        marginTop:20,
        float :'right',
        
    }
}))
interface Iprops {

}
export function AddConfig(props: Iprops) {
    const classes = useStyles()
    const [openForm, setOpenForm] = React.useState(false)
    const { register, handleSubmit } = useForm();
    const [state, setState] = useState({})
    const [services,setServices]=useState([])

    const onSubmit = (data: any) => {
         data['services']=services
         console.log(data, 'all data confic')
        
    }
    return (
        <div className={classes.root}>
        <h2> add new config</h2>
            <form  onSubmit={handleSubmit(onSubmit)}>
            <Grid container spacing={3} >
              <Grid item xs={4}>
                 <MasTextField   inputRef1={register} label={'host'} name={'host'} required={false} />
              </Grid>
              <Grid item xs={4}>
                 <MasTextField  inputRef1={register} label={'http_port'} name={'http_port'} required={false} />
              </Grid>
              <Grid item xs={4}>
                 <MasTextField  inputRef1={register} label={'tls_port'} name={'tls_port'} required={false} />
              </Grid>
              <Grid item xs={4}>
                 <MasTextField  inputRef1={register} label={'max_call_recv_msg_size'} name={'max_call_recv_msg_size'} required={false} />
              </Grid>
              <Grid item xs={4}>
                  <MasTextField  inputRef1={register} label={'websocket_ping_interval'} name={'websocket_ping_interval'} required={false} />
              </Grid>
              <Grid item xs={4}>
                   <MasTextField  inputRef1={register} label={'http_max_write_timeout'} name={'http_max_write_timeout'} required={false} />
              </Grid>
              <Grid item xs={4}>
                  <MasTextField  inputRef1={register} label={'http_max_read_timeout'} name={'http_max_read_timeout'} required={false} />
              </Grid>              
             
            </Grid>
            <Grid container xs={12} spacing={2}>
              <Grid item xs={3}>
                  <MasCheckBox name='run_http_server' label='run_http_server' inputRef1={register}/>
              </Grid>
                        
              <Grid item xs={3}>
                  <MasCheckBox name='run_tls_server' label='run_tls_server' inputRef1={register}/>
              </Grid>
              <Grid item xs={3}>
                  <MasCheckBox name='allow_all_origin' label='allow_all_origin' inputRef1={register}/>
              </Grid>
              <Grid item xs={3}>
                  <MasCheckBox name='use_websockets' label='use_websockets' inputRef1={register} onClick={handleSubmit(onSubmit)}/>
              </Grid>
              </Grid>
        <div className={classes.btn}>
            <MasButton type="submit" label='save' variant='outlined' color='secondary'/>
            <MasButton label='add New Servise' variant='outlined' type="button" color='primary' onClick={()=>setOpenForm(true)}/>
        </div>
            </form>
                <AddService openForm={openForm} onSave={(data)=>{
                    setServices(data)
                    setOpenForm(false)
                }} onClose={()=>{setOpenForm(false)}} />

        </div>
    )
}
