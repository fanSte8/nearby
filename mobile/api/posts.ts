import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

export const getPosts = async (longitude: number, latitude: number) => {
  try {
    const response = await axios.get(URLS.POSTS.GET_ALL, { params: { longitude, latitude } })
    return response.data;
  } catch(error) {
    return formatError(error)
  }
}

export const likePost = async (id: number): Promise<boolean> => {
  try {
    await axios.post(URLS.POSTS.LIKE.replace(':id', String(id)))
    return true
  } catch(error) {
    return false
  }
}
