import axios from "axios";
import { URLS } from "./urls";
import * as SecureStore from 'expo-secure-store';
import { JWT_KEY } from "../constants";

interface Response {
  data: any,
  error: any
}

export const login = async (email: string, password: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.LOGIN, { email, password })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: error.response.data.error
    }
  }
}

export const register = async (firstName: string, lastName: string, email: string, password: string) => {
  axios.post(URLS.USERS.REGISTER, { firstName, lastName, email, password })
  .then(res => console.log(res.data))
  .catch(err => console.log(err.data.error));
}
