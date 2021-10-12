import React from "react";
import clsx from "clsx";
import { makeStyles, TextField, TextFieldProps } from "@material-ui/core";

export interface IMasTextField {
  onChange?: any;
  onBlur?: any;
  name?: string;
  value?: any;
  type?: any;
  defaultValue?: any;
  labelStyle?: any;
  styleTextFiled?: any;
  required?: boolean;
  autoFocus?: any;
  label: string;
  placeholder?: string | any;
  size?: "medium" | "small";
  variant?: "filled" | "outlined" | "standard";
  color?: "primary" | "secondary";
  inputRef1?: any;
  error?: boolean;
  disabled?: boolean;
  id?: string;
  helperText?: string;
  multiline?: boolean;
  rows?: string;
  rowsMax?: string;
  displayMemper?: string;
  InputProps?: any
  inputProps?: any
}
const useStyles = makeStyles(() => ({
  root: {
    border: "1px solid #ddd",
    marginBottom: "12px",
  },
}));

export function MasTextField(_props: IMasTextField) {
  const classes = useStyles();
  const {
    onChange,
    onBlur,
    label,
    name,
    inputRef1,
    autoFocus,
    error,
    styleTextFiled,
    disabled,
    value,
    defaultValue,
    variant,
    color,
    size,
    required,
    id,
    helperText,
    type,
    multiline,
    rows,
    rowsMax,
    placeholder,
    InputProps,
    inputProps,
    ...restProps
  } = _props;
  return (
    <TextField  inputRef={inputRef1} label={label}
    name={name}
      InputLabelProps={{ shrink: true }}
      variant={variant || "filled"}
      color={color}
      placeholder={placeholder}
      id={id}
      size={size}
      disabled={disabled}
      defaultValue={defaultValue}
      value={value}
      className={clsx(classes.root, [styleTextFiled])}
      helperText={helperText}
      type={type}
      fullWidth
      rows={rows}
      rowsMax={rowsMax}
      multiline={multiline}
      onChange={onChange}
      InputProps={InputProps}
      inputProps={inputProps}
    />
  );
}
