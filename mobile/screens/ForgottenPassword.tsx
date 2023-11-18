import { useState } from "react";
import { View, StyleSheet, Text } from "react-native";
import { LabeledInput,Button } from "../components";

export const ForgottenPasswordScreen = () => {
  const [email, setEmail] = useState('');

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Forgot your password? Enter your email and we'll send you a code you can use to reset it.</Text>
      <View style={styles.inputContainer}>
        <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
      </View>
      <View style={styles.buttons}>
        <Button onPress={() => null} text="Send email" />
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: '#ffffff',
    width: "100%",
    padding: 10
  },
  logo: {
    width: 100,
    height: 100,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  inputContainer: {
    width: "100%",
    padding: 10
  },
  buttons: {
    alignSelf: 'center'
  },
  text: {
    paddingHorizontal: 20
  }
});
