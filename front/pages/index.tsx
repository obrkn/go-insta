import type { NextPage } from 'next';
import Head from 'next/head';
// import Image from 'next/image';
import { Container, Box, Button } from '@mui/material';
import axios from 'axios';
import { useState, useCallback } from 'react';

const Home: NextPage = () => {
  const [data, setData] = useState("まだ何もデータ来てないよう")
  const getData = useCallback(
    () => {
      axios
        .get("http://localhost:8080")
        .then(res => setData(res.data))
        // .get("back:8080", { 
        //   headers: {'withCredentials': 'true'}
        // })
        // .then(res => setData(res.data))
    },
    [],
  )
  return (
    <Container maxWidth="sm">
      <Head>
        <title>tiwitter</title>
        <meta name="description" content="Goを使ったインスタ風アプリです。" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <Box sx={{ m: 2 }}>
          {data}
        </Box>
        <Button onClick={() => getData()}>押して！</Button>
      </main>

      <footer>
      　あああああ
      </footer>
    </Container>
  )
}

export default Home