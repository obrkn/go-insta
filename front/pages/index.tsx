import React from 'react';
import type { NextPage } from 'next';
import Head from 'next/head';
import Link from 'next/link'
// import Image from 'next/image';
import { Container, Box, Button } from '@mui/material';
import { ApiWithToken, Api } from '../components/axios';
import { useState, useCallback } from 'react';
import Context from '../lib/store/context'

const themes = {
  light: {
    foreground: "#000000",
    background: "#eeeeee"
  },
  dark: {
    foreground: "#ffffff",
    background: "#222222"
  }
};

const Home: NextPage = () => {
  const [data, setData] = useState("まだ何もデータ来てないよう")
  const [text, setText] = useState("")
  const postData = useCallback(async(event) => {
    event.preventDefault();
    ApiWithToken.post("/post", {text}).then(res => console.log(res.data))
  }, [text])
      const { state, dispatch } = React.useContext(Context)
  return (
    
    <Container maxWidth="sm">
      <Head>
        <title>tiwitter</title>
        <meta name="description" content="Goを使ったインスタ風アプリです。" />
        <link rel="icon" href="/favicon.ico" />
      </Head>


 <button onClick={() => dispatch({ type: 'SIGNIN' })}>+</button>
 <p>{state.signedIn.toString()}</p>
      <main>
        <Box sx={{ m: 2 }}>
          {data}
        </Box>
        <Button onClick={() => console.log(process.env.NEXT_PUBLIC_API_BASE_URL)}>押して！</Button>
        <form onSubmit={postData}>
          <label>
            Name:
            <input type="text" value={text} onChange={e => setText(e.target.value)} />
          </label>
          <input type="submit" value="Submit" />
        </form>
      </main>

      <Link href={'/signup'}>
      サインアップ
      </Link>
      <footer>
      　あああああ
      </footer>
    </Container>
  )
}

export default Home