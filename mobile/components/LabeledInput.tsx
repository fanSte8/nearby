import { View, Text, KeyboardTypeOptions } from "react-native"
import { Input } from "./Input"

interface PropsType {
  label: string,
  value: string,
  onChangeText: (text: string) => void,
  placeholder: string,
  secureText?: boolean,
  keyboardType?: KeyboardTypeOptions
  multiline?: boolean
  style?: any
}

export const LabeledInput = ({ label, value, onChangeText, placeholder, secureText, keyboardType, multiline, style }: PropsType) => {
  return <View style={{ padding: 10 }}>
    <Text style={{ paddingLeft: 5 }}>{label}</Text>
    <Input 
      value={value}
      onChangeText={onChangeText}
      placeholder={placeholder}
      secureText={secureText || false}
      keyboardType={keyboardType || undefined}
      multiline={multiline}
      style={style}
    />
  </View>
}
