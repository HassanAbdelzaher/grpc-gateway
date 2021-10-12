import React from 'react';
import RadioGroup from '@material-ui/core/RadioGroup';
import { FormControl, FormLabel, FormControlLabel, Radio } from '@material-ui/core'


interface props {
    formlabel: string,
    itemText: string,
    disabled?: boolean,
    items: Array<any>,
    itemValue: string,
    inputRef: any,
    name: string
    stleFormControl?: string,
    styleLabel?: string
}
export function MasRadioGroup(_props: props) {
    const { formlabel, stleFormControl, styleLabel, itemText, disabled, items, itemValue, inputRef, name } = _props

    return (
        <FormControl className={stleFormControl} >
            <FormLabel component="legend" className={styleLabel}>{formlabel}</FormLabel>
            <RadioGroup name={name}>
                {items.map((item: any) => {
                    return <FormControlLabel inputRef={inputRef} value={item[itemValue]} disabled={disabled} control={<Radio />} label={item[itemText]} />
                })}
            </RadioGroup>
        </FormControl>
    );
}