import React, { useState, useEffect } from 'react'
import { View, Text, Image, TouchableOpacity, Alert } from 'react-native'
import { useUserStore } from '../storage/useUserStorage'
import { SafeAreaView } from 'react-native-safe-area-context'
import { Ionicons } from '@expo/vector-icons'
import { PRIMARY_COLOR } from '../constants'
import { getUserById, uploadAvatar } from '../api/users'
import * as ImagePicker from 'expo-image-picker'

export const AccountScreen = ({ navigation, route }: any) => {
  const user = useUserStore(state => state.user)
  const [avatar, setAvatar] = useState<string>('')
  const [firstName, setFirstName] = useState<string>('')
  const [lastName, setLastName] = useState<string>('')
  const [email, setEmail] = useState<string>('')

  useEffect(() => {
    (async () => {
      if (!route.params.id) {
        return
      }

      const user = await getUserById(route.params.id)
      if (!user) {
        return
      }

      setFirstName(user.firstName)
      setLastName(user.lastName)
      setEmail(user.email)
      setAvatar(user.imageUrl)
    })()
  }, [])

  const isCurrentUser = user && user.id === route.params.id

  const handleAvatarChange = async () => {
    if (isCurrentUser) {
      let result = await ImagePicker.launchImageLibraryAsync({
        mediaTypes: ImagePicker.MediaTypeOptions.Images,
        allowsEditing: true,
        aspect: [1, 1],
        quality: 1,
      })

      if (!result.canceled) {
        try {
          const imageUrl = await uploadAvatar(result.assets[0].uri)
          if (imageUrl) {
            setAvatar(imageUrl)
          } else {
            Alert.alert('Error', 'Failed to upload image')
          }
        } catch (error) {
          console.error('Error uploading image:', error)
          Alert.alert('Error', 'Failed to upload image')
        }
      }
    }
  }

  return (
    <SafeAreaView>
      <TouchableOpacity onPress={navigation.goBack} style={{ alignSelf: 'flex-start' }}>
        <Ionicons name="chevron-back" color={"black"} size={32}/>
      </TouchableOpacity>
      <View style={{ alignItems: 'center', marginTop: 20 }}>
        <Image source={avatar ? { uri: avatar } : require('../assets/default-avatar.png')} style={{ width: 150, height: 150, borderRadius: 75 }} />
        {isCurrentUser && (
          <TouchableOpacity onPress={handleAvatarChange}>
            <Text style={{ color: PRIMARY_COLOR, marginTop: 10 }}>Change avatar</Text>
          </TouchableOpacity>
        )}
        <Text style={{ fontSize: 24, marginTop: 10 }}>{`${user?.firstName} ${user?.lastName}`}</Text>
        <Text style={{ fontSize: 18, marginTop: 5 }}>{user?.email}</Text>
      </View>
    </SafeAreaView>
  )
}
