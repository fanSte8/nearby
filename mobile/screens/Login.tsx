import { useState } from "react"
import { View, Image, Text, StyleSheet, TouchableOpacity } from "react-native"
import { LabeledInput, Button, Alert } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { AuthLayout } from "../layouts"
import { login } from "../api/users"
import { useUserStore } from "../storage/useUserStorage"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"

export const LoginScreen = ({ navigation, route }: any) => {
  const fromScreen = route?.params?.from;

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const setUser = useUserStore(store => store.setUser)
  const setToken = useUserStore(store => store.setToken)
  const setIsLoggedIn = useUserStore(store => store.setIsLoggedIn)

  const handleLogin = async () => {
    const { data, error } = await login(email, password)
    
    if (error) {
      setError(error)
    } else {
      setToken(data.token)
      setUser(data.user)
      setIsLoggedIn(true)
    }
  }

  return (
    <AuthLayout>
      <View style={styles.container}>
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>
          <View style={{ padding: 10 }}>
            {
              error && (
                <Alert type='warning' text="Incorrect email or password" />
              )
            }
            {
              fromScreen === 'register' && (
                <Alert type='success' text="Your account has been created. An activation code has been sent to your email. You can activate your account in the setting once you login." />
              )
            }
            {
              fromScreen === 'reset-password' && (
                <Alert type='success' text="Your password has been reset." />
              )
            }
          </View>
          <View style={styles.inputContainer}>
            <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
            <LabeledInput label="Password" value={password} onChangeText={setPassword} placeholder="" secureText={true} />
          </View>
          <View style={styles.buttons}>
          <Button onPress={handleLogin} text="Login" />
            <TouchableOpacity style={styles.link} onPress={() => navigation.navigate('ForgottenPassword')}>
              <Text style={styles.linkText}>Forgotten password?</Text>
            </TouchableOpacity>
            <TouchableOpacity style={styles.link} onPress={() => navigation.navigate('Register')}>
              <Text style={styles.linkText}>Don't have an account yet? Register!</Text>
            </TouchableOpacity>
          </View>
        </KeyboardAwareScrollView>
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
})
