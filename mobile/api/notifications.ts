import { URLS } from "./urls"
import { formatError, authorizedAxios as axios } from "."

export const hasSeenNotifications = async () => {
  try {
    const response = await axios.get(URLS.NOTIFICATIONS.HAS_UNSEEN)
    return response.data.seen
  } catch(error) {
    return false
  }
}