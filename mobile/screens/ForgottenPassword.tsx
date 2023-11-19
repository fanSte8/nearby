import { useState } from "react"
import { View, StyleSheet, Text, TouchableOpacity } from "react-native"
import { LabeledInput,Button, Alert } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { AuthLayout } from "../layouts"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"
import { forgottenPassword } from "../api/users"

export const ForgottenPasswordScreen = ({ navigation }: any) => {
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')

  const handleForgottenPassword = async () => {
    const { error } = await forgottenPassword(email)

    if (error) {
      setError(error)
    } else {
      navigation.navigate('ResetPassword', { email })
    }
  }

  return (
    <AuthLayout>
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>
          <View style={styles.container}>
            <Text style={styles.text}>Forgot your password? Enter your email and we'll send you a code you can use to reset it.</Text>
            {
              error && (
                <Alert type='warning' text={error} />
              )
            }
            <View style={styles.inputContainer}>
              <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
            </View>
            <View style={styles.buttons}>
              <Button onPress={handleForgottenPassword} text="Send email" />
              <TouchableOpacity style={styles.link}  onPress={() => navigation.navigate('Login')}>
                <Text style={styles.linkText}>Back to Login screen</Text>
              </TouchableOpacity>
            </View>
          </View>
        </KeyboardAwareScrollView>
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
  text: {
    paddingHorizontal: 20
  },
  link: {
    alignSelf: 'center',
    paddingTop: 10
  },
  linkText: {
    color: PRIMARY_COLOR
  }
})
