import { useState } from "react"
import { View, Text, StyleSheet, TouchableOpacity, ScrollView } from "react-native"
import { LabeledInput, Button, Alert } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { NearbyLogoLayout } from "../layouts"
import { register } from "../api/users"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"
import { useUserStore } from "../storage/useUserStorage"
import { useNavigation } from "@react-navigation/native"

export const RegisterScreen = ({ navigation }: any) => {
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [error, setError] = useState('')
  
  const setUser = useUserStore(store => store.setUser)
  const setToken = useUserStore(store => store.setToken)
  const setIsLoggedIn = useUserStore(store => store.setIsLoggedIn)

  const handleRegister = async () => {
    if (password !== confirmPassword) {
      setError("Passwords do not match")
      return
    }

    const { data, error } = await register(firstName, lastName, email, password)


    if (error) {
      setError(error)
    } else {
      setToken(data.token)
      setUser(data.user)
      setIsLoggedIn(true)
    }
  }

  return (
    <NearbyLogoLayout navigation={navigation}>
      <ScrollView style={styles.container}>
        {
          error && (
            <Alert type='warning' text={error} />
          )
        }
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>
          <View style={styles.inputContainer}>
            <LabeledInput value={firstName} onChangeText={setFirstName} label="First Name" placeholder="" secureText={false} />
            <LabeledInput value={lastName} onChangeText={setLastName} label="Last Name" placeholder="" secureText={false} />
            <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
            <LabeledInput label="Password" value={password} onChangeText={setPassword} placeholder="" secureText={true} />
            <LabeledInput label="Confirm Password" value={confirmPassword} onChangeText={setConfirmPassword} placeholder="" secureText={true} />
          </View>
          <View style={styles.buttons}>
            <Button onPress={handleRegister} text="Register" />
            <TouchableOpacity style={styles.link} onPress={() => navigation.navigate('Login')}>
              <Text style={styles.linkText}>Already have an account? Login here!</Text>
            </TouchableOpacity>
          </View>
        </KeyboardAwareScrollView>
      </ScrollView>
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
    height: 100
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
