import { useState } from "react"
import { View, Image, Text, StyleSheet, TouchableOpacity } from "react-native"
import { LabeledInput, Button } from "../components";
import { PRIMARY_COLOR } from "../constants";
import { useNavigation } from "@react-navigation/native";
import { AuthLayout } from "../layouts";

export const LoginScreen = ({ navigation }: any) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <AuthLayout>
      <View style={styles.container}>
        <View style={styles.inputContainer}>
          <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
          <LabeledInput label="Password" value={password} onChangeText={setPassword} placeholder="" secureText={true} />
        </View>
        <View style={styles.buttons}>
        <Button onPress={() => null} text="Login" />
          <TouchableOpacity style={styles.link} onPress={() => navigation.navigate('ForgottenPassword')}>
            <Text style={styles.linkText}>Forgotten password?</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.link} onPress={() => navigation.navigate('Register')}>
            <Text style={styles.linkText}>Don't have an account yet? Register!</Text>
          </TouchableOpacity>
        </View>
      </View>
    </AuthLayout>
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
  link: {
    alignSelf: 'center',
    paddingTop: 10
  },
  linkText: {
    color: PRIMARY_COLOR
  }
});
