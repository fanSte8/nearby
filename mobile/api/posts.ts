import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

export const getPosts = async (sort: string, longitude: number, latitude: number, page: number, pageSize = 10) => {
  try {
    const response = await axios.get(URLS.POSTS.GET_ALL, { params: { longitude, latitude, sort, page, pageSize } })
    return response.data
  } catch(error) {
    return formatError(error)
  }
}

export const likePost = async (id: number): Promise<boolean> => {
  try {
    const res = await axios.post(URLS.POSTS.LIKES.replace(':id', String(id)))
    return res.data.like
    return true
  } catch(error) {
    return false
  }
}

export const getComments = async (postId: number, page = 1, pageSize = 10) => {
  try {
    const response = await axios.get(URLS.POSTS.COMMENTS.replace(':id', String(postId)), { params: { page, pageSize } })
    return response.data
  } catch(error) {
    return []
  }
}

export const postComment = async (postId: number, text: string) => {
  try {
    const response = await axios.post(URLS.POSTS.COMMENTS.replace(':id', String(postId)), { text })
    return response.data
  } catch(error) {
    return null
  }
}
