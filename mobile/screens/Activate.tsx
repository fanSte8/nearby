import { useState } from "react";
import { Alert, Button, Input, LabeledInput } from "../components"
import { NearbyLogoLayout } from "../layouts"
import { StyleSheet, TouchableOpacity, Text, View } from "react-native";
import { PRIMARY_COLOR } from "../constants";
import { useNavigation } from "@react-navigation/native";
import { activateAccount, sendActivationCode } from "../api/users";

export const ActivateScreen = ({ navigation }: any) => {
  const [code, setCode] = useState('')
  const [error, setError] = useState('')
  const [alert, setAlert] = useState('')

  const handleActivateAccount = async () => {
    const { error } = await activateAccount(code)

    if (error) {
      setError(error)
    } else {
      setError('')
      setAlert('Your account has been activated. You need to log out for the changes to take effect.')
    }
  }

  return (
    <NearbyLogoLayout navigation={navigation}>
      <View style={styles.container}>
        <Alert type='warning' text={error} />
        <Alert type='success' text={alert} />
        <Text style={styles.text}>An activation code has been sent to your email. Enter it here to activate your account</Text>
        <Input value={code} onChangeText={setCode} placeholder="code" />
        <Button onPress={handleActivateAccount} text="Activate" />
        <TouchableOpacity style={styles.link}  onPress={sendActivationCode}>
          <Text style={styles.linkText}>Didn't receive an email? Click here to send again</Text>
        </TouchableOpacity>
      </View>
    </NearbyLogoLayout>
  )
}

const styles = StyleSheet.create({
  text: {
    padding: 10,
    textAlign: "center"
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
