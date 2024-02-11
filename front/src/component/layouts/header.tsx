import {
  AccountCircle,
  Link as LinkIcon,
  Mail,
  Notifications,
  ReceiptLong,
} from "@mui/icons-material";
import {
  AppBar,
  Badge,
  Box,
  IconButton,
  Toolbar,
  Typography,
} from "@mui/material";
import Link from "next/link";
import * as React from "react";

export default function Header() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Link href="/">
            <Box sx={{ display: "flex", gap: 1, alignItems: "center" }}>
              <LinkIcon />
              <Typography variant="h6" noWrap component="div">
                Short URL
              </Typography>
            </Box>
          </Link>

          <Box sx={{ flexGrow: 1 }} />
          <Box>
            <IconButton size="large" edge="end" color="inherit" href="/history">
              <ReceiptLong />
            </IconButton>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
