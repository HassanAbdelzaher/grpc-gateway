import React from "react";
import {FormControl, InputLabel, OutlinedInput} from "@material-ui/core";
interface props {
  styleLabel?: any;
  styleInput?: any;
  styleFormControl?: any;
  error?: any;
  className?: any;
  autoComplete?: string;
  labelWidth?: number;
  required?: boolean;
  type?: string;
  id?: string;
  label?: string;
  inputRef?: any;
  name?: string;
  endAdornment?: any;
  value?: string;
  multiline?: boolean;
  onChange?: any;
  placeholder?: any;
  defaultValue?:any
}
export function MasInput(props: props) {
  const {
    styleLabel,
    styleInput,
    styleFormControl,
    error,
    className,
    endAdornment,
    autoComplete,
    labelWidth,
    required,
    type,
    id,
    label,
    inputRef,
    name,
    value,
    multiline,
    onChange,
    placeholder,
    defaultValue,
    ...InputProps
  } = props;
  return (
    <FormControl required={required} className={styleFormControl}>
      <InputLabel className={styleLabel} htmlFor={id}>
        {label}
      </InputLabel>
      <OutlinedInput
        id={id}
        name={name}
        className={styleInput}
        autoComplete={autoComplete}
        inputRef={inputRef}
        type={type}
        error={error}
        multiline={multiline}
        labelWidth={labelWidth}
        endAdornment={endAdornment}
        value={value}
        placeholder={placeholder}
        onChange={onChange}
        defaultValue={defaultValue}
        {...InputProps}
      />
    </FormControl>
  );
}
