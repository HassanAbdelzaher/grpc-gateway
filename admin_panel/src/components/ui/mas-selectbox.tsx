/* eslint-disable no-use-before-define */
import React, {useEffect} from "react";
import Autocomplete from "@material-ui/lab/Autocomplete";
import {makeStyles, Theme, createStyles} from "@material-ui/core/styles";
import {IMasCombo} from "./mas-compobox";
import {MasTextField} from ".";

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
      minHeight: "150PX",
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
interface ISelectBox extends IMasCombo {}

export function MasSelectBox(props: ISelectBox) {
  const {
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
  const [value, setValue] = React.useState(null);
  const classes = useStyles();
  useEffect(() => {
        setValue(defaultValue);
  }, [defaultValue]);

  return (
    <Autocomplete
      id={id}
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
      disableClearable
      //disableCloseOnSelect
      getOptionLabel={(option) => (option ? option[displayMemper] : "")}
      onChange={(event, newValue) => {
        setValue(newValue);
        if (onChange) {
          onChange(event, newValue);
        }
      }}
      renderOption={(option) => {
       return  <React.Fragment>
          <div className={classes.text}>
            {option ? option[displayMemper] : ""}
          </div>
        </React.Fragment>
      }}
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
