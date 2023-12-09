import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

export const getPosts = async (sort: string, longitude: number, latitude: number, page: number, pageSize = 10) => {
  try {
    const response = await axios.get(URLS.POSTS.POSTS, { params: { longitude, latitude, sort, page, pageSize } })
    return response.data
  } catch(error) {
    return formatError(error)
  }
}

export const getUserPosts = async (userId: number, longitude: number, latitude: number, page: number, pageSize = 10) => {
  try {
    const response = await axios.get(URLS.POSTS.POSTS.replace(':id', String(userId)), { params: { longitude, latitude, page, pageSize } })
    return response.data
  } catch(error) {
    return formatError(error)
  }
}

export const likePost = async (id: number): Promise<boolean> => {
  try {
    const res = await axios.post(URLS.POSTS.LIKES.replace(':id', String(id)))
    return res.data.like
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


export const createPost = async (description: string, photo: string, latitude: string, longitude: string) => {
  try {
    const formData = new FormData()
    formData.append('image', {
      uri: photo,
      name: 'photo.jpg',
      type: 'image/jpg',
    } as any)
    formData.append('description', description)
    formData.append('latitude', latitude)
    formData.append('longitude', longitude)

    await axios.post(URLS.POSTS.POSTS, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })

    return true
  } catch (error) {
    return false
  }
}

export const getPostById = async (id: number) => {
  try {
    const response = await axios.get(URLS.POSTS.POST_BY_ID.replace(':id', String(id)))
    return response.data
  } catch(error) {
    return null
  }
}
