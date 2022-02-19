import axios from 'axios';

export const Api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  timeout: 5000,
  withCredentials: true,
})

export const ApiWithToken = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  timeout: 5000,
  withCredentials: true,
})

ApiWithToken.interceptors.request.use( async config => {
  const resp = await Api.get('/token')
  const token = resp.headers["x-csrf-token"]
  if (token && config.headers) {
    config.headers['x-csrf-token'] = token
  }
  return config
})