import { useState } from "react"
import { View, StyleSheet, Text, TouchableOpacity } from "react-native"
import { LabeledInput,Button, Alert } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { NearbyLogoLayout } from "../layouts"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"
import { forgottenPassword, resetPassword } from "../api/users"

export const ResetPasswordScreen = ({ navigation, route }: any) => {
  const email = route?.params?.email

  const [code, setCode] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleResetPassword = async () => {
    const { error } = await resetPassword(password, code)

    if (error) {
      setError(error)
    } else {
      navigation.navigate('Login', { from: 'reset-password' })
    }
  }

  return (
    <NearbyLogoLayout navigation={navigation}>
      <View style={styles.container}>
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>
            {
              email && (
                <Alert type='success' text={`A password reset code has been sent to your email: ${email}.`} />
              )
            }
            {
              error && (
                <Alert type='warning' text={error} />
              )
            }
          <View style={styles.inputContainer}>
            <LabeledInput value={code} onChangeText={setCode} label="Code" placeholder="" secureText={false} />
            <LabeledInput value={password} onChangeText={setPassword} label="New Password" placeholder="" secureText={true} />
          </View>
          <View style={styles.buttons}>
            <Button onPress={handleResetPassword} text="Reset password" />
            <TouchableOpacity style={styles.link}  onPress={() => forgottenPassword(email)}>
              <Text style={styles.linkText}>Didn't receive an email? Click here to try again.</Text>
            </TouchableOpacity>
            <TouchableOpacity style={styles.link}  onPress={() => navigation.navigate('ForgottenPassword')}>
              <Text style={styles.linkText}>Not your email? Click here to change it</Text>
            </TouchableOpacity>
            <TouchableOpacity style={styles.link}  onPress={() => navigation.navigate('Login')}>
              <Text style={styles.linkText}>Back to Login screen</Text>
            </TouchableOpacity>
          </View>
        </KeyboardAwareScrollView>
      </View>
    </NearbyLogoLayout>
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
  },
  text: {
    paddingHorizontal: 20
  }
})
