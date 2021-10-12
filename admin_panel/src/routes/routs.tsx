import React from 'react'
import { Route, Switch, Redirect } from 'react-router-dom';
import {AddConfig} from '../pages/add-config'
import {AddService} from '../pages/add-service'
import {Main} from '../pages/main'
const NotFound=()=>{
    return <div style={{color:'red'}}> not found</div>
}
export default function Routes() {
    return (
        <>
            <Switch >
                <Route exact path="/" component={Main}></Route>
                <Route  path="/addconfig" component={AddConfig}></Route>
                <Route  path="/addService" component={AddService}></Route>
                <Route path='/notfound' component={NotFound}></Route>
                <Redirect to="/notfound" />
            </Switch>
        </>
    )
}
