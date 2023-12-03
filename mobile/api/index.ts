import axios from 'axios'
import { useUserStore } from '../storage/useUserStorage'

export const formatError = (error: any) => {
  if (typeof error.response.data.error === 'object') {
    return Object.values(error.response.data.error).join('\n')
  }

  return error.response.data.error
}

export const authorizedAxios = axios.create()

authorizedAxios.interceptors.request.use(
  (config) => {
    const token = useUserStore.getState().token
    config.headers['Authorization'] = `Bearer ${token}`
  
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

authorizedAxios.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      useUserStore.getState().reset()
    }
    return Promise.reject(error)
  }
)
