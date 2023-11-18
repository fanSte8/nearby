import { StyleSheet, Text, View } from 'react-native';
import { ForgottenPasswordScreen, LoginScreen, RegisterScreen } from './screens';
import { AuthLayout } from './layouts';

const App = () => {
  return (
    <View style={styles.container}>
      <AuthLayout>
        <RegisterScreen />
      </AuthLayout>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-start',
  },
});

export default App;
