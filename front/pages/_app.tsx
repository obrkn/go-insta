import React from 'react'
import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Context from '../lib/store/context'
import { reducer } from '../lib/store/reducer'

function MyApp({ Component, pageProps }: AppProps) {
  const initialState = { signedIn: false }
  const [state, dispatch] = React.useReducer(reducer, initialState)

  return (
    <Context.Provider value={{ state, dispatch }}>
      <Component {...pageProps} />
    </Context.Provider>
  )
}

export default MyApp
