"use client";

import { Api, LinkResponse } from "@/lib/api";
import { saveClipboard } from "@/service/save-clipboard";
import { useEffect, useState } from "react";
import UrlList from "@/component/molecules/url-list";

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

  return <UrlList urlList={urlList} />;
}
