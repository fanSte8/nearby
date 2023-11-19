import { StyleSheet, Text, View } from 'react-native';
import { ForgottenPasswordScreen, LoginScreen, RegisterScreen, ResetPasswordScreen } from './screens';
import { AuthLayout } from './layouts';

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

const Stack = createNativeStackNavigator();

const App = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName='Login' screenOptions={{ headerShown: false }}>
        <Stack.Screen name='Login' component={LoginScreen} />
        <Stack.Screen name='Register' component={RegisterScreen} />
        <Stack.Screen name='ForgottenPassword' component={ForgottenPasswordScreen} />
        <Stack.Screen name='ResetPassword' component={ResetPasswordScreen} />
      </Stack.Navigator>
  </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-start',
  },
});

export default App;
