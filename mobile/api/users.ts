import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

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
      error: formatError(error)
    }
  }
}

export const register = async (firstName: string, lastName: string, email: string, password: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.REGISTER, { firstName, lastName, email, password })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const forgottenPassword = async (email: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.FORGOTTEN_PASSWORD, { email })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    if (error.response.status === 404) {
      return {
        data: null,
        error: 'No user found with the provided email.'
      }
    }

    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const resetPassword = async (password: string, code: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.RESET_PASSWORD, { password, code })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const sendActivationCode = async () => {
  axios.get(URLS.USERS.ACTIVATE)
}

export const activateAccount = async (token: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.ACTIVATE, { token })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const changePassword = async (oldPassword: string, password: string): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.CHANGE_PASSWORD, { password, oldPassword })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const changeRadius = async (radius: number): Promise<Response> => {
  try {
    const response = await axios.post(URLS.USERS.CHANGE_RADIUS, { radius })

    return {
      data: response.data,
      error: null
    }
  } catch(error: any) {
    return {
      data: null,
      error: formatError(error)
    }
  }
}

export const uploadAvatar = async (imageUri: string): Promise<string> => {
  const formData = new FormData()
  formData.append('profile-picture', {
    uri: imageUri,
    name: `avatar_${Date.now()}`,
    type: 'image/jpeg',
  } as any)

  try {
    const response = await axios.post(URLS.USERS.PROFILE_PICTURE, formData, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'multipart/form-data'
      },
    })

    if (response.status === 200) {
      return response.data.imageUrl
    }
    return ''
  } catch (error) {
    return ''
  }
}

export const getUserById = async (id: number) => {
  try {
    const response = await axios.get(URLS.USERS.USER_BY_ID.replace(':id', id.toString()))
    return response.data.user
  } catch(error) {
    return null
  }
}
