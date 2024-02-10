"use client"

import { TextField, TextFieldProps } from "@mui/material"

export default function RoundedInput(props: TextFieldProps) {
  return (
    <TextField
      {...props}
      fullWidth
      sx={{
        "& .MuiInputBase-root": {
          borderRadius: "3rem",
        },
        "& .MuiInputBase-input": {
          px: "2rem",
          fontSize: "1.5rem",
        },
      }}
    />
  )
}
