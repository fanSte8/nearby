import { TouchableOpacity, Text, StyleSheet } from "react-native"
import { PRIMARY_COLOR } from "../constants"

interface PropsType {
  onPress: () => void,
  text: string
}

export const Button = ({ onPress, text }: PropsType) => {
  return <TouchableOpacity onPress={onPress} style={styles.button}>
    <Text style={styles.buttonText}>{text}</Text>
  </TouchableOpacity>
}

const styles = StyleSheet.create({
  button: {
    backgroundColor: PRIMARY_COLOR,
    paddingVertical: 10,
    paddingHorizontal: 20,
    borderRadius: 8,
    alignItems: 'center',
    marginVertical: 10
  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: 'bold',
  },
})