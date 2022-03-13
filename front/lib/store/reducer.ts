import { State, Action } from '../interfaces'

export const reducer = (state: State, action: Action) => {
  switch (action.type) {
    case 'SIGNIN':
      return {
        ...state,
        signedIn: true,
      }
    case 'SIGNOUT':
      return {
        ...state,
        signedIn: false,
      }
    default:
      return state
  }
}