import { View, Text } from "react-native"
import { Input } from "./Input";

interface PropsType {
  label: string,
  value: string,
  onChangeText: (text: string) => void,
  placeholder: string,
  secureText: boolean
}

export const LabeledInput = ({ label, value, onChangeText, placeholder, secureText }: PropsType) => {
  return <View style={{ padding: 10 }}>
    <Text style={{ paddingLeft: 5 }}>{label}</Text>
    <Input 
      value={value}
      onChangeText={onChangeText}
      placeholder={placeholder}
      secureText={secureText}
    />
  </View>
}
