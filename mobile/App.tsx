import { StyleSheet, Text, View } from 'react-native'
import { ForgottenPasswordScreen, HomeScreen, LocationPermissionDeniedScreen, LoginScreen, RegisterScreen, ResetPasswordScreen } from './screens'
import * as Location from 'expo-location';  

import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { useUserStore } from './storage/useUserStorage'
import { useEffect, useState } from 'react';

const Stack = createNativeStackNavigator()

const App = () => {
  const isLoggedIn = useUserStore(store => store.isLoggedIn)
  const [locationPermission, setLocationPermission] = useState(false)

  useEffect(() => {
    (async () => {
      let { status } = await Location.requestForegroundPermissionsAsync();
      setLocationPermission(status === 'granted')
    })();
  });

  if (!locationPermission) {
    return <LocationPermissionDeniedScreen />
  }

  return (
    <NavigationContainer>
      {
        isLoggedIn ? (
          <Stack.Navigator initialRouteName='Home' screenOptions={{ headerShown: false }}>
            <Stack.Screen name='Home' component={HomeScreen} />
          </Stack.Navigator>
        ) : (
          <Stack.Navigator initialRouteName='Login' screenOptions={{ headerShown: false }}>
            <Stack.Screen name='Login' component={LoginScreen} />
            <Stack.Screen name='Register' component={RegisterScreen} />
            <Stack.Screen name='ForgottenPassword' component={ForgottenPasswordScreen} />
            <Stack.Screen name='ResetPassword' component={ResetPasswordScreen} />
          </Stack.Navigator>
        )
      }
    </NavigationContainer>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-start',
  },
})

export default App
