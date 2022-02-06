import type { NextPage } from 'next';
import Head from 'next/head';
// import Image from 'next/image';
import { Container, Box, Button } from '@mui/material';
import axios from 'axios';
import { useState, useCallback } from 'react';

const Home: NextPage = () => {
  const [data, setData] = useState("まだ何もデータ来てないよう")
  const [text, setText] = useState("")
  const getData = useCallback(
    () => {
      axios
        .get("http://localhost:8080/api/home")
        .then(res => setData(res.data))
        // .get("http://localhost:8080/api/home")
        // .then(res => setData(res.data))
        // .get("back:8080", { 
        //   headers: {'withCredentials': 'true'}
        // })
        // .then(res => setData(res.data))
    },
    [],
  )
  const postData = useCallback(async(event) => {
    event.preventDefault();
    const instance_d = axios.create({
      withCredentials: true,
    })
    const resp = await instance_d.get("http://localhost:8080/api/token")
    const token = resp.headers["x-csrf-token"]
    const instance = axios.create({
      withCredentials: true,
      headers: {"X-CSRF-Token": token}
    })
    instance.post("http://localhost:8080/api/post", {text})
  }, [text])
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
        <form onSubmit={postData}>
          <label>
            Name:
            <input type="text" value={text} onChange={e => setText(e.target.value)} />
          </label>
          <input type="submit" value="Submit" />
        </form>
      </main>

      <footer>
      　あああああ
      </footer>
    </Container>
  )
}

export default Home