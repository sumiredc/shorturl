"use client";

import { Api, LinkResponse } from "@/lib/api";
import { Link } from "@mui/icons-material";
import {
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { useEffect, useState } from "react";

export default function HistoryPage() {
  const [urlList, setUrlList] = useState<LinkResponse[]>([]);
  let init = false;

  async function getUrlList() {
    const api = new Api();
    try {
      const { data } = await api.url.getUrl();
      setUrlList(data);
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    if (init) {
      return;
    }

    getUrlList();

    return () => {
      init = true;
    };
  }, []);

  return (
    <List>
      {urlList.map((url, idx) => (
        <ListItem key={idx} disablePadding>
          <ListItemButton href={url.short} target="_blank">
            <ListItemIcon>
              <Link />
            </ListItemIcon>
            <ListItemText primary={url.short} />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
}
