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

    if (!url) {
      setLoading(false);
      setErrorMessage("入力が必要です");
      return;
    }
    setErrorMessage("");

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
