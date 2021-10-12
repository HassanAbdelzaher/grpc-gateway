import React from "react";
import Select, {SelectProps} from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormControl from "@material-ui/core/FormControl";
import {makeStyles} from "@material-ui/core";
import clsx from "clsx";
// import NativeSelect from "@material-ui/core/NativeSelect";

const useStyles = makeStyles(() => ({
  formControl: {
    width: "100%",
  },
  lable: {},
  menu: {
    flexGrow: 1,
    border: "1px solid #ddd",
    borderRadius: "0",
  },
  menuItem: {},
}));

export interface MasSelectProps extends SelectProps {
  items: any[];
  itemText: string | null;
  itemValue: string | null;
  label: string;
  labelStyle?: string;
  disabled?: any;
  defaultValue?: any;
  onChange?: (e: any) => void;
  styleItem?: string;
  styleMenue?: string;
  styleFormControl?: string;
  name?: string;
  multiple?: boolean;
}
export function MasSelect(ownProp: MasSelectProps) {
  const classes = useStyles();
  const {
    items,
    onChange,
    multiple,
    styleItem,
    styleMenue,
    styleFormControl,
    inputRef,
    disabled,
    name,
    defaultValue,
    labelStyle,
    label,
    itemValue,
    itemText,
    ...restProps
  } = ownProp;
  const MasSelectitems: any[] = items;

  return (
    <FormControl
      variant="filled"
      className={clsx(styleFormControl, classes.formControl)}
    >
      <InputLabel
        shrink={true}
        className={clsx(labelStyle, classes.lable)}
        htmlFor={name}
      >
        {label}
      </InputLabel>
      <Select
        native
        name={name}
        defaultValue={defaultValue}
        disabled={disabled}
        className={clsx(styleMenue, classes.menu)}
        multiple={multiple}
        onChange={(e) => (onChange ? onChange(e) : null)}
        inputProps={{
          id: name,
        }}
        {...restProps}
      >
        <option aria-label="None" value="" />
        {MasSelectitems ? (
          MasSelectitems.map((item: any, index: number) => {
            return (
              <option
                key={index} 
                value={
                  itemValue ? (item[itemValue] ? item[itemValue] : null) : index
                }
              >
                {itemText ? item[itemText] : item}
              </option>
            );
          })
        ) : (
          <MenuItem className={classes.menuItem} key={1} value={0} />
        )}
      </Select>
    </FormControl>
  );
}
