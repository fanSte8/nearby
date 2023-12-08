import { TextInput, View, Text, StyleSheet, KeyboardTypeOptions } from "react-native"

interface PropsType {
  value: string,
  onChangeText: (text: string) => void,
  placeholder: string,
  secureText?: boolean
  keyboardType?: KeyboardTypeOptions
}

export const Input = ({ value, onChangeText, placeholder, secureText, keyboardType }: PropsType) => {
  return <TextInput 
    style={styles.inputStyle}
    value={value}
    onChangeText={onChangeText}
    placeholder={placeholder}
    secureTextEntry={secureText || false}
    keyboardType={keyboardType || undefined}
  />
}

const styles = StyleSheet.create({
  inputStyle: {
    width: "100%",
    borderWidth: 1,
    borderColor: '#ccc',
    padding: 8,
    fontSize: 16,
    borderRadius: 8
  }
})