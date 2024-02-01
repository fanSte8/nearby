import { LogBox, StyleSheet } from 'react-native'
import {
  AccountScreen,
  ActivateScreen,
  ChangePasswordScreen,
  ChangeRadius as ChangeRadiusScreen,
  CreatePostScreen,
  ForgottenPasswordScreen,
  HomeScreen,
  LocationPermissionDeniedScreen,
  LoginScreen,
  NotificationsScreen,
  PostDetails as PostDetailsScreen,
  RegisterScreen,
  ResetPasswordScreen
} from './screens'
import * as Location from 'expo-location'  

LogBox.ignoreAllLogs(true)

import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { useUserStore } from './storage/useUserStorage'
import { useEffect, useState } from 'react'

const Stack = createNativeStackNavigator()

const App = () => {
  const isLoggedIn = useUserStore(store => store.isLoggedIn)
  const [locationPermission, setLocationPermission] = useState(false)

  useEffect(() => {
    (async () => {
      let { status } = await Location.requestForegroundPermissionsAsync()
      setLocationPermission(status === 'granted')
    })()
  })

  if (!locationPermission) {
    return <LocationPermissionDeniedScreen />
  }

  return (
    <NavigationContainer>
      {
        isLoggedIn ? (
          <Stack.Navigator initialRouteName='Home' screenOptions={{ headerShown: false }}>
            <Stack.Screen name='Home' component={HomeScreen} />
            <Stack.Screen name="PostDetails" component={PostDetailsScreen} />
            <Stack.Screen name="Activate" component={ActivateScreen} />
            <Stack.Screen name="ChangePassword" component={ChangePasswordScreen} />
            <Stack.Screen name="ChangeRadius" component={ChangeRadiusScreen} />
            <Stack.Screen name='Account' component={AccountScreen} />
            <Stack.Screen name='Notifications' component={NotificationsScreen} />
            <Stack.Screen name='CreatePost' component={CreatePostScreen} />
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
