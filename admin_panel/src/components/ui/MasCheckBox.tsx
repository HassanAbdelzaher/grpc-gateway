import React from 'react';
import {Checkbox,FormControl,FormControlLabel,CheckboxProps } from '@material-ui/core';

export interface props extends CheckboxProps{
   name:string
   inputRef1?:any
   label:string
   labelPlacement?:"bottom"|"end"|"start"|"top"
   styleFormControl?:string
   styleCheckbox?:string
}
export function MasCheckBox(_props: props) {
    const { inputRef1,checked,color,styleFormControl,styleCheckbox,defaultChecked,labelPlacement,label, name,onChange } = _props
    return (
    <FormControl className={styleFormControl}>
       <FormControlLabel
        label={label}
        name={name}
        inputRef={inputRef1}
        labelPlacement={labelPlacement}
          control={<Checkbox
                    className={styleCheckbox}
                    //  name={name}
                    //  inputRef={inputRef1}
                     color={color}
                     defaultChecked={defaultChecked}  
                     onChange={onChange}
                     checked={checked}
                />}
         
        />
        </FormControl>

    );
}