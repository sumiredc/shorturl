import UrlForm from "@/component/organisms/url-form";
import { Box, Container } from "@mui/material";

export default function Home() {
  return (
    <Container component="main">
      <Box sx={{ my: 10 }}>
        <UrlForm />
      </Box>
    </Container>
  );
}
