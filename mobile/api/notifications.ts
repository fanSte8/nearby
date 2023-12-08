import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

export const hasSeenNotifications = async () => {
  try {
    const response = await axios.get(URLS.NOTIFICATIONS.SEEN)
    return response.data.seen
  } catch(error) {
    return false
  }
}

export const getNotifications = async (page: number, pageSize: number) => {
  try {
    const response = await axios.get(URLS.NOTIFICATIONS.NOTIFICATIONS, { params: { page, pageSize } })
    return response.data.notifications
  } catch (error) {
    return []
  }
}
