import { useState } from "react"
import { View, StyleSheet, Text, TouchableOpacity } from "react-native"
import { LabeledInput,Button } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { AuthLayout } from "../layouts"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"

export const ResetPasswordScreen = ({ navigation }: any) => {
  const [code, setCode] = useState('')
  const [password, setPassword] = useState('')

  return (
    <AuthLayout>
      <View style={styles.container}>
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>
          <Text style={styles.text}>Forgot your password? Enter your email and we'll send you a code you can use to reset it.</Text>
          <View style={styles.inputContainer}>
            <LabeledInput value={code} onChangeText={setCode} label="Code" placeholder="" secureText={false} />
            <LabeledInput value={password} onChangeText={setPassword} label="New Password" placeholder="" secureText={true} />
          </View>
          <View style={styles.buttons}>
            <Button onPress={() => null} text="Reset password" />
            <TouchableOpacity style={styles.link}  onPress={() => navigation.navigate('Login')}>
              <Text style={styles.linkText}>Didn't receive an email? Click here to send it again.</Text>
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
  },
  text: {
    paddingHorizontal: 20
  }
})
