import {
  AccountCircle,
  Link as LinkIcon,
  Mail,
  Notifications,
} from "@mui/icons-material"
import {
  AppBar,
  Badge,
  Box,
  IconButton,
  Toolbar,
  Typography,
} from "@mui/material"
import Link from "next/link"
import * as React from "react"

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
            <IconButton
              size="large"
              aria-label="show 4 new mails"
              color="inherit"
            >
              <Badge badgeContent={4} color="error">
                <Mail />
              </Badge>
            </IconButton>
            <IconButton size="large" color="inherit">
              <Badge badgeContent={17} color="error">
                <Notifications />
              </Badge>
            </IconButton>
            <IconButton size="large" edge="end" color="inherit">
              <AccountCircle />
            </IconButton>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  )
}
