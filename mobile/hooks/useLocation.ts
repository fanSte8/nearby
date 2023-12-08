import { useState, useEffect } from "react"
import * as Location from 'expo-location'

export const useLocation = () => {
  const [latitude, setLatitude] = useState(0)
  const [longitude, setLongitude] = useState(0)

  useEffect(() => {
    (async () => {
      const location = await Location.getCurrentPositionAsync({})
      setLatitude(location.coords.latitude)
      setLongitude(location.coords.longitude)
    })()
  }, [])

  return { latitude, longitude }
} 