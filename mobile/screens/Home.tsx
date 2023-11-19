import { View, Text, TouchableOpacity } from "react-native"
import { useUserStore } from "../storage/useUserStorage"

export const HomeScreen = () => {
  const user = useUserStore(store => store.user)
  const reset = useUserStore(store => store.reset)

  return (
    <View>
      <Text>{JSON.stringify(user)}</Text>
      <TouchableOpacity onPress={reset}><Text>logout</Text></TouchableOpacity>
    </View>
  )
}
