import { saveClipboard } from "@/service/save-clipboard";
import { LinkResponse } from "@/lib/api";
import { Link } from "@mui/icons-material";
import {
  Alert,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Snackbar,
} from "@mui/material";
import { useState } from "react";

type Props = {
  urlList: LinkResponse[];
};

export default function UrlList({ urlList }: Props) {
  const [snackbar, setSnackbar] = useState(false);

  function handleClick(url: LinkResponse) {
    saveClipboard(url.short);
    setSnackbar(true);
  }

  function handleClose() {
    setSnackbar(false);
  }

  return (
    <>
      <List>
        {urlList.length ? (
          urlList.map((url, idx) => (
            <ListItem key={idx} disablePadding>
              <ListItemButton onClick={() => handleClick(url)}>
                <ListItemIcon>
                  <Link />
                </ListItemIcon>
                <ListItemText primary={url.short} secondary={url.original} />
              </ListItemButton>
            </ListItem>
          ))
        ) : (
          <Alert severity="info" sx={{ my: 3 }}>
            履歴が有りません
          </Alert>
        )}
      </List>
      <Snackbar
        open={snackbar}
        onClose={handleClose}
        autoHideDuration={1200}
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <Alert severity="success" sx={{ width: "100%" }}>
          コピーしました
        </Alert>
      </Snackbar>
    </>
  );
}
