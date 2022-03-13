import * as React from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import {
  Avatar,
  Button,
  CssBaseline,
  TextField,
  Link,
  Box,
  Typography,
  Container,
  Alert,
  Modal,
} from '@mui/material';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { ApiWithToken, Api } from '../components/axios';
import '../components/axios.ts';

const theme = createTheme();

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 2,
};

const SignUp: NextPage = () => {
  const [error, setError] = React.useState('');
  const [modalVisible, setModalVisible] = React.useState(false);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const params = new URLSearchParams(new FormData(event.currentTarget) as any);
    ApiWithToken.post('/signup', params)
      .then(res => setModalVisible(true))
      .catch(err => setError(err.response?.data || '新規登録に失敗しました。'))
  };

  return (
  <>
    {error && <Alert severity="error">{error}</Alert>}
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
              id="nickname"
              label="Nickname"
              name="nickname"
              autoComplete="nickname"
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
    <Modal
      open={modalVisible}
      onClose={() => setModalVisible(false)}
    >
      <Box sx={style}>
        <Typography sx={{textAlign: 'center'}}>
          サインアップが完了しました。
        </Typography>
        <Button href='/signin' variant="contained" sx={{mt: 2, width: '100%'}}>
          OK
        </Button>
      </Box>
    </Modal>
  </>
  );
}

export default SignUp