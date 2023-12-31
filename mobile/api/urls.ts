export const URLS = {
  USERS: {
    LOGIN: `${process.env.EXPO_PUBLIC_BASE_URL}/users/login`,
    REGISTER: `${process.env.EXPO_PUBLIC_BASE_URL}/users/register`,
    FORGOTTEN_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/forgotten-password`,
    RESET_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/reset-password`,
    CHANGE_PASSWORD: `${process.env.EXPO_PUBLIC_BASE_URL}/users/change-password`,
    CHANGE_RADIUS: `${process.env.EXPO_PUBLIC_BASE_URL}/users/posts-radius`,
    ACTIVATE: `${process.env.EXPO_PUBLIC_BASE_URL}/users/activate`,
    PROFILE_PICTURE: `${process.env.EXPO_PUBLIC_BASE_URL}/users/profile-picture`,
    USER_BY_ID: `${process.env.EXPO_PUBLIC_BASE_URL}/users/:id`
  },
  POSTS: {
    POST_BY_ID: `${process.env.EXPO_PUBLIC_BASE_URL}/posts/:id`,
    POSTS: `${process.env.EXPO_PUBLIC_BASE_URL}/posts`,
    LIKES: `${process.env.EXPO_PUBLIC_BASE_URL}/posts/:id/likes`,
    COMMENTS: `${process.env.EXPO_PUBLIC_BASE_URL}/posts/:id/comments`,
    USER_POSTS: `${process.env.EXPO_PUBLIC_BASE_URL}/posts/users/:id`
  },
  NOTIFICATIONS: {
    NOTIFICATIONS: `${process.env.EXPO_PUBLIC_BASE_URL}/notifications`,
    SEEN: `${process.env.EXPO_PUBLIC_BASE_URL}/notifications/seen`
  }
}