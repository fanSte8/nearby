import React, { useState } from 'react'
import { View, Alert, Image, TouchableOpacity, StyleSheet } from 'react-native'
import * as ImagePicker from 'expo-image-picker'
import { MaterialIcons } from '@expo/vector-icons'
import { useLocation } from '../hooks/useLocation'
import { createPost } from '../api/posts'
import { PRIMARY_COLOR } from '../constants'
import Ionicons from '@expo/vector-icons/Ionicons'
import { Button, Input, LabeledInput } from '../components'
import { SafeAreaView } from 'react-native-safe-area-context'

export const CreatePostScreen = ({ navigation }: any) => {
  const [description, setDescription] = useState('')
  const [photo, setPhoto] = useState('')
  const [width, setWidth] = useState(1)
  const [height, setHeight] = useState(1)
  
  const { latitude, longitude } = useLocation()

  const takePhoto = async () => {
    let permissionResult = await ImagePicker.requestCameraPermissionsAsync()
    if (permissionResult.granted === false) {
      Alert.alert('Permission to access camera is required!')
      return
    }

    let pickerResult = await ImagePicker.launchCameraAsync({})

    if (!pickerResult.canceled) {
      setPhoto(pickerResult.assets[0].uri)

      setWidth(pickerResult.assets[0].width)
      setHeight(pickerResult.assets[0].height)
    }
  }

  const handlerCreatePost = async () => {
    if (!description) {
      Alert.alert("Description is required")
      return
    }

    if (!photo) {
      Alert.alert("No photo has been taken yet.")
      return
    }

    const success = await createPost(description, photo, String(latitude), String(longitude))
    if (!success) {
      Alert.alert('Unable to create post')
    }

    navigation.navigate('Home')
  }

  return (
    <SafeAreaView style={styles.container}>
      <TouchableOpacity onPress={navigation.goBack} style={{ alignSelf: 'flex-start' }}>
        <Ionicons name="chevron-back" color={"black"} size={32}/>
      </TouchableOpacity>
      <Input
        value={description}
        onChangeText={setDescription}
        placeholder="Enter description"
        multiline={true}
        style={{
          borderWidth: 0,
          borderBottomWidth: 1
        }}
      />
      {photo && <Image
          source={{ uri: photo }}
          style={{
            width: '100%',
            maxHeight: 500,
            marginTop: 20,
            aspectRatio: width / height,
            resizeMode: 'contain',
          }}/>}
      <View style={styles.buttonContainer}>
        <TouchableOpacity onPress={takePhoto}>
          <MaterialIcons name="add-a-photo" size={32} color={PRIMARY_COLOR} style={styles.takePhotoButton} />
        </TouchableOpacity>
        <Button text="Create Post" onPress={handlerCreatePost} />
      </View>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 10,
    justifyContent: 'flex-start',
    alignItems: 'center',
  },
  descriptionInput: {
    width: '100%',
    height: 100,
    marginBottom: 20,
    borderWidth: 1,
    borderColor: 'gray',
    padding: 10,
  },
  buttonContainer: {
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignItems: 'flex-end',
    width: '100%',
    position: 'absolute',
    bottom: 20,
    paddingHorizontal: 0,
  },
  takePhotoButton: {
    marginRight: 10
  },
})
