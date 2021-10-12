/* eslint-disable no-use-before-define */
import React, {useEffect} from "react";
import Autocomplete from "@material-ui/lab/Autocomplete";
import {MasTextField} from ".";
import {IMasTextField} from "./MasTextField";
import {makeStyles, Theme, createStyles} from "@material-ui/core/styles";
import DoneIcon from "@material-ui/icons/Done";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    paper: {
      boxShadow: "none",
      margin: 0,
      color: "#586069",
      fontSize: 13,
      position: "relative",
    },
    option: {
      minHeight: "auto",
      alignItems: "flex-start",
      padding: 8,
      '&[aria-selected="true"]': {
        backgroundColor: "transparent",
      },
      '&[data-focus="true"]': {
        backgroundColor: theme.palette.action.hover,
      },
    },
    popperDisablePortal: {
      position: "relative",
    },
    popper: {
      position: "absolute",
      minHeight: "150px",
    },
    iconSelected: {
      width: 15,
      height: 15,
      marginRight: 5,
      marginLeft: 3,
      color: "#fff",
      backgroundColor: theme.palette.primary.main,
    },
    color: {
      width: 14,
      height: 14,
      flexShrink: 0,
      borderRadius: 3,
      marginRight: 8,
      marginTop: 2,
    },
    text: {
      flexGrow: 1,
      color: "#282b2e",
      fontSize: "1rem",
    },
    close: {
      opacity: 0.6,
      width: 18,
      height: 18,
    },
    input: {
      marginBottom: 0,
    },
  })
);
export interface IMasCombo extends IMasTextField {
  options: Array<any>;
  onChange?: (event:any, value:any) => void;
  required: boolean;
  className?: string;
  displayMemper: string;
  ValueMember?: any;
  defaultValue?: any;
  disablePortal?: boolean;
  tags?: number;
}
export function MasCompoBox(props: IMasCombo) {
  const {
    tags,
    className,
    options,
    displayMemper,
    onChange,
    defaultValue,
    label,
    placeholder,
    id,
    variant,
    required,
    styleTextFiled
  } = props;
  const [value, setValue] = React.useState<Array<any>>([]);
  const classes = useStyles();
  useEffect(() => {
    if (defaultValue) {
      setValue(defaultValue);
    }
  }, [defaultValue]);

  return (
    <Autocomplete
      id={id}
      multiple
      limitTags={tags || 3}
      classes={{
        paper: classes.paper,
        option: classes.option,
        popper: classes.popper,
        popperDisablePortal: classes.popperDisablePortal,
      }}
      className={className}
      options={options}
      selectOnFocus={false}
      value={value}
      onChange={(event, newValue) => {
        setValue(newValue);
        if (onChange) {
          onChange(event, newValue);
        }
      }}
      disableClearable
      disableCloseOnSelect
     // getOptionSelected={(option, value) => option.NAME === value.NAME}
      getOptionLabel={(option) => (option ? option[displayMemper] : "")}
      renderOption={(option, {selected}) => (
        <React.Fragment>
          <DoneIcon
            className={classes.iconSelected}
            style={{visibility: selected ? "visible" : "hidden"}}
          />
          <span
            className={classes.color}
            style={{backgroundColor: option.color}}
          />
          <div className={classes.text}>
            {option ? option[displayMemper] : ""}
          </div>
        </React.Fragment>
      )}
      renderInput={(params) => (
        <MasTextField
          {...params}
          required={required}
          variant={variant}
          label={label}
          styleTextFiled={styleTextFiled}
          placeholder={placeholder}
        />
      )}
    />
  );
}
