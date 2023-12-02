import { useState } from "react"
import { NearbyLogoLayout } from "../layouts"
import { Alert, Button, LabeledInput } from "../components"
import { View, StyleSheet, Text } from "react-native"
import { PRIMARY_COLOR } from "../constants"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"
import { changePassword } from "../api/users"

export const ChangePasswordScreen = ({ navigation }: any) => {
  const [oldPassword, setOldPassword] = useState('')
  const [newPassword, setNewPassword] = useState('')
  const [confirmNewPassword, setConfirmNewPassword] = useState('')
  const [error, setError] = useState('')
  const [alert, setAlert] = useState('')

  const handleChangePassword = async () => {
    if (newPassword !== confirmNewPassword) {
      setError('Passwords do not match')
      return
    }

    const { error } = await changePassword(oldPassword, newPassword)
    if (error) {
      setAlert('')
      setError(error)
    } else {
      setError('')
      setAlert('Password changed successfully')
    }
  }

  return (
    <NearbyLogoLayout navigation={navigation}>
      <View style={styles.container}>
        <Text style={styles.text}>Change Password</Text>
        {
          error && <Alert type='warning' text={error} />
        }
        {
          alert && <Alert type='success' text={alert} />
        }
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>  
          <LabeledInput value={oldPassword} onChangeText={setOldPassword} label="Old Password" placeholder="" secureText={true} />
          <LabeledInput value={newPassword} onChangeText={setNewPassword} label="New Password" placeholder="" secureText={true} />
          <LabeledInput value={confirmNewPassword} onChangeText={setConfirmNewPassword} label="Confirm Password" placeholder="" secureText={true} />
          <Button onPress={handleChangePassword} text="Change Password" />
        </KeyboardAwareScrollView>
      </View>
    </NearbyLogoLayout>
  )
}

const styles = StyleSheet.create({
  text: {
    padding: 10,
    textAlign: "center",
    fontSize: 18
  },
  container: {
    backgroundColor: '#ffffff',
    width: "100%",
    height: "100%",
    padding: 10
  },
  link: {
    alignSelf: 'center',
    paddingTop: 10
  },
  linkText: {
    color: PRIMARY_COLOR
  }
});
