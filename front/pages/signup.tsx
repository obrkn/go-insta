import * as React from 'react';
import type { NextPage } from 'next';
import {
  Avatar,
  Button,
  CssBaseline,
  TextField,
  Link,
  Box,
  Typography,
  Container,
} from '@mui/material';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { ApiWithToken, Api } from '../components/axios';
import '../components/axios.ts';

const theme = createTheme();

const SignUp: NextPage = () => {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const params = new URLSearchParams(new FormData(event.currentTarget) as any);
    ApiWithToken.post('/signup', params)
      .then(res => console.log(`成功：${res}`))
      .catch(err => console.log(`失敗：${err}`))
  };

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Avatar alt="Cindy Baker" src="/animal_chara_computer_penguin.png" />
          <Typography component="h1" variant="h5">
            Twitter風アプリ
          </Typography>
          <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              サインアップ
            </Button>
            <Box sx={{justifyContent: 'center'}}>
              <Link href="/signin">
                既にアカウントをお持ちの方はこちら
              </Link>
            </Box>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}

export default SignUp