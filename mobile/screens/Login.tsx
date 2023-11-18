import { useState } from "react"
import { View, Image, Text, StyleSheet, TouchableOpacity } from "react-native"
import { LabeledInput, Button } from "../components";
import { PRIMARY_COLOR } from "../constants";

export const LoginScreen = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <View style={styles.container}>
      <View style={styles.inputContainer}>
        <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
        <LabeledInput label="Password" value={password} onChangeText={setPassword} placeholder="" secureText={true} />
      </View>
      <Button onPress={() => null} text="Login" />
      <TouchableOpacity style={styles.link}>
        <Text style={styles.linkText}>Forgotten password?</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.link}>
        <Text style={styles.linkText}>Don't have an account yet? Register!</Text>
      </TouchableOpacity>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
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
  link: {
    paddingTop: 10
  },
  linkText: {
    color: PRIMARY_COLOR
  }
});
