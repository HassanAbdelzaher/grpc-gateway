import React from 'react';
import Switch from '@material-ui/core/Switch';
import {FormControl,FormControlLabel } from '@material-ui/core';

 interface props{
    name:string
    inputRef:any
    checked?:boolean
    defaultChecked?:boolean
    label:string
    labelPlacement?:"bottom"|"end"|"start"|"top"
    color?:"primary"|"secondary"|"default";
    styleFormControl?:string
    styleSwich?:string
  }
export function MasSwitch(_props: props) {
    const { inputRef,color,styleFormControl,styleSwich,defaultChecked,labelPlacement,label, name, ...restProps } = _props

    return (
    <FormControl className={styleFormControl}>
       <FormControlLabel
          label={label}
          labelPlacement={labelPlacement}
          control={
            <Switch
               className={styleSwich}
               name={name}
               inputRef={inputRef}
               color={color}
               defaultChecked={defaultChecked}
              {...restProps }
            />
          }
        />
        </FormControl>
    );
}