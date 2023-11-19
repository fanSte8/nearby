import { useState } from "react";
import { View, StyleSheet, Text, TouchableOpacity } from "react-native";
import { LabeledInput,Button } from "../components";
import { PRIMARY_COLOR } from "../constants";
import { AuthLayout } from "../layouts";

export const ForgottenPasswordScreen = ({ navigation }: any) => {
  const [email, setEmail] = useState('');

  return (
    <AuthLayout>
      <View style={styles.container}>
        <Text style={styles.text}>Forgot your password? Enter your email and we'll send you a code you can use to reset it.</Text>
        <View style={styles.inputContainer}>
          <LabeledInput value={email} onChangeText={setEmail} label="Email" placeholder="" secureText={false} />
        </View>
        <View style={styles.buttons}>
          <Button onPress={() =>  navigation.navigate('ResetPassword')} text="Send email" />
          <TouchableOpacity style={styles.link}  onPress={() => navigation.navigate('Login')}>
            <Text style={styles.linkText}>Back to Login screen</Text>
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
});
