import { LinkResponse } from "@/lib/api";
import { Link } from "@mui/icons-material";
import {
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";

type Props = {
  urlList: LinkResponse[];
};

export default function UrlList({ urlList }: Props) {
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
