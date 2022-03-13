export type State = {
  signedIn: boolean,
}

export type Action = { type: 'SIGNIN' } | { type: 'SIGNOUT' }