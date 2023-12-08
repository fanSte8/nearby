import { SafeAreaView, TouchableOpacity, Text } from "react-native"
import { Ionicons } from '@expo/vector-icons'

export const AccountScreen = ({ navigation }: any) => {
  return (
    <SafeAreaView>
      <TouchableOpacity onPress={navigation.goBack} style={{ alignSelf: 'flex-start' }}>
        <Ionicons name="chevron-back" color={"black"} size={32}/>
      </TouchableOpacity>
      <Text>Account screen</Text>
    </SafeAreaView>
  )
}