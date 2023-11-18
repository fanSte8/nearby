import { TextInput, View, Text, StyleSheet } from "react-native"

interface PropsType {
  value: string,
  onChangeText: (text: string) => void,
  placeholder: string,
  secureText: boolean
}

export const Input = ({ value, onChangeText, placeholder, secureText }: PropsType) => {
  return <TextInput 
    style={styles.inputStyle}
    value={value}
    onChangeText={onChangeText}
    placeholder={placeholder}
    secureTextEntry={secureText}
  />
}

const styles = StyleSheet.create({
  inputStyle: {
    width: "100%",
    borderWidth: 1,
    borderColor: '#ccc',
    padding: 8,
    fontSize: 16,
  }
});