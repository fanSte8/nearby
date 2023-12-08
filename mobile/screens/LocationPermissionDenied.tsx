import { View, Text } from "react-native"
import { NearbyLogoLayout } from "../layouts"

export const LocationPermissionDeniedScreen = () => {
  return (
    <NearbyLogoLayout>
      <Text style={{ fontWeight: 'bold', fontSize: 20, padding: 40 }}>
        Location access has been denied. In order to use the app, access to location data is required.
      </Text>
    </NearbyLogoLayout>
  )
}
