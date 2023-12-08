import { SafeAreaView, TouchableOpacity, Text } from "react-native"
import { Ionicons } from '@expo/vector-icons'

export const CreatePostScreen = ({ navigation }: any) => {
  return (
    <SafeAreaView>
      <TouchableOpacity onPress={navigation.goBack} style={{ alignSelf: 'flex-start' }}>
        <Ionicons name="chevron-back" color={"black"} size={32}/>
      </TouchableOpacity>
      <Text>Create post screen</Text>
    </SafeAreaView>
  )
}