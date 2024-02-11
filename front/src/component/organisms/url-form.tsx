"use client";

import {
  CircularProgress,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemText,
} from "@mui/material";
import RoundedInput from "../atoms/form/rounded-input";
import { Send } from "@mui/icons-material";
import { ChangeEvent, useState } from "react";
import { Api, LinkResponse } from "@/lib/api";
import UrlList from "../molecules/url-list";
import { UrlStoreValidator } from "@/service/validator/url-store-validator";
import { ValidationException } from "@/exceptions/validation-exception";

export default function UrlForm() {
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [urlList, setUrlList] = useState<LinkResponse[]>([]);

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    setUrl(e.target.value);
  }

  function handleClick() {
    setLoading(true);
    setErrorMessage("");

    try {
      const validator = new UrlStoreValidator(url);
      validator.validate();
    } catch (e: unknown) {
      if (e instanceof ValidationException) {
        setErrorMessage(e.message);
      } else {
        console.error(e);
      }
      setLoading(false);
      return;
    }

    createUrl().finally(() => {
      setLoading(false);
    });
  }

  async function createUrl() {
    const api = new Api();
    try {
      const { data } = await api.url.postUrl({ original: url });
      setUrlList([...urlList, data]);
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <>
      <RoundedInput
        value={url}
        disabled={loading}
        error={!!errorMessage}
        helperText={errorMessage}
        onChange={handleChange}
        InputProps={{
          endAdornment: (
            <IconButton disabled={loading} onClick={handleClick}>
              <Send />
              {loading && (
                <CircularProgress
                  size={40}
                  sx={{
                    position: "absolute",
                    top: 0,
                    left: -2,
                    zIndex: 1,
                  }}
                />
              )}
            </IconButton>
          ),
        }}
      />
      <UrlList urlList={urlList} />
    </>
  );
}
