import React, { useState } from 'react'
import { makeStyles, Grid } from '@material-ui/core'
import { MasTextField, MasButton, MasCheckBox, MasSelect } from '../components/ui'
import { useForm } from "react-hook-form";

const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: "80%",
        margin: " 5px auto",
        padding: '5px',
        boxShadow: "0, 0, 30px rgba(0,0,0,0.4)",
        background: 'white'
    },
    styleFormControl: {
        width: '100%'
    }
}))
interface Iprops {

}
export function AddService(props: Iprops) {
    const classes = useStyles()
    const { register, handleSubmit } = useForm();
    const [state, setState] = useState<{ service_name: string, backends: Array<any> }[]>([])
    const [serviceName, setServiceName] = useState('')

    const onSubmit = (data: any) => {
        if(serviceName=='')
        alert('choose service name')
         return;
        let temp = {}
        let obj: { service_name: string, backends: Array<any> } = {
            service_name: serviceName,
            backends: []
        }
        for (let k in data) {
            temp[k] = data[k]
        }
        obj.backends.push(temp)
        state.push(obj)
        setState([...state])
    }
    console.log(state)
    const handleChange = (e) => {
        setServiceName(e.target.value)
    }
    return (
        <div className={classes.root}>
            <h1> add backends</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <Grid item xs={12}>
                    <MasSelect onChange={handleChange} items={[{ name: 'service1' }, { name: 'service2' }]} itemText='name' itemValue='name' label={'service_name'} name={'service_name'} required={false} />
                </Grid>
                <h3>backend configration</h3><br />
                <Grid container spacing={3} >
                    <Grid item xs={4}>
                        <MasTextField inputRef1={register} label={'max_call_recv_msg_size'} name={'max_call_recv_msg_size'} required={false} type='number' />
                    </Grid>
                    <Grid item xs={4}>
                        <MasTextField inputRef1={register} label={'tls_client_cert'} name={'tls_client_cert'} required={false} />
                    </Grid>
                    <Grid item xs={4}>
                        <MasTextField inputRef1={register} label={'tls_client_key'} name={'tls_client_key'} required={false} />
                    </Grid>
                    <Grid item xs={4}>
                        <MasTextField inputRef1={register} label={'backoff_max_delay'} name={'backoff_max_delay'} required={false} />
                    </Grid>
                    <Grid item xs={4}>
                        <MasCheckBox name='tls_no_verify' label='tls_no_verify' inputRef1={register} />
                    </Grid>
                    <Grid item xs={4}>
                        <MasCheckBox name='is_using_tls' label='is_using_tls' inputRef1={register} />
                    </Grid>
                    <Grid item xs={4}>
                        <MasTextField inputRef1={register} label={'default_authority'} name={'default_authority'} required={false} />
                    </Grid>
                </Grid>
                <MasButton type="submit" label='save' variant='contained' color='secondary' />
            </form>

        </div>
    )
}
