export const URLS = {
  USERS: {
    LOGIN: `${process.env.EXPO_PUBLIC_BASE_URL}/users/login`,
    REGISTER: `${process.env.EXPO_PUBLIC_BASE_URL}/users/register`,
    FORGOTTEN_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/forgotten-password`,
    RESET_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/reset-password`,
    CHANGE_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/change-password`,
    CHANGE_RADIUS: `${process.env.EXPO_PUBLIC_BASE_URL}/users/posts-radius`,
    ACTIVATE: `${process.env.EXPO_PUBLIC_BASE_URL}/users/activate`,
  },
  POSTS: {
    GET_ALL: `${process.env.EXPO_PUBLIC_BASE_URL}/posts`,
    LIKE: `${process.env.EXPO_PUBLIC_BASE_URL}/posts/:id/like`
  }
}