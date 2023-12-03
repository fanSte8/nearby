import moment from "moment"

export const formatDistance = (distance: number) => {
  return `${Math.round(distance / 100) / 10}km away`
}

export const formatTime = (time: string) => {
  return moment.utc(time).local().startOf('seconds').fromNow()
}
