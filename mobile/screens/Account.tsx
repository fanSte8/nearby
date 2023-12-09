import React, { useState, useEffect } from 'react'
import { View, Text, Image, TouchableOpacity, Alert, FlatList, ScrollView } from 'react-native'
import { useUserStore } from '../storage/useUserStorage'
import * as Location from 'expo-location'
import { PRIMARY_COLOR } from '../constants'
import { getUserById, uploadAvatar } from '../api/users'
import * as ImagePicker from 'expo-image-picker'
import { Loading, Post } from '../components'
import { BasicLayout } from '../layouts'
import { getUserPosts } from '../api/posts'
import { usePostsStore } from '../storage/usePostsStorage'

export const AccountScreen = ({ navigation, route }: any) => {
  const user = useUserStore(state => state.user)
  const posts = usePostsStore(state => state.posts)
  const addPosts = usePostsStore(state => state.addPosts)
  const reset = usePostsStore(state => state.reset)

  const [avatar, setAvatar] = useState<string>('')
  const [firstName, setFirstName] = useState<string>('')
  const [lastName, setLastName] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    (async () => {
      reset()
      await fetchPosts()
    })()
  }, [])

  useEffect(() => {
    (async () => {
      if (!route.params.id) {
        return
      }

      setLoading(true)
      const user = await getUserById(route.params.id)
      if (!user) {
        return
      }

      setFirstName(user.firstName)
      setLastName(user.lastName)
      setEmail(user.email)
      setAvatar(user.imageUrl)
      setLoading(false)
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
  
  const pageSize = 20
  const [page, setPage] = useState(1)
  const [hasMorePosts, setHasMorePosts] = useState(true)
  const [isLoadingPosts, setIsLoadingPosts] = useState(false)

  const fetchPosts = async (nextPage = page) => {
    if (!hasMorePosts || isLoadingPosts) {
      return
    }

    const location = await Location.getCurrentPositionAsync({})

    if (user) {
      setIsLoadingPosts(true)
      const res = await getUserPosts(user.id, location.coords.longitude, location.coords.latitude, nextPage, pageSize)
      const newPosts = res?.posts || []

      if (newPosts.length < pageSize) {
        setHasMorePosts(false)
      }
      
      addPosts(newPosts)
      setPage(nextPage + 1)
      setIsLoadingPosts(false)
    }
  }

  if (loading) {
    return <Loading />
  }

  return (
    <BasicLayout navigation={navigation} title={"User"}>
      <ScrollView showsVerticalScrollIndicator={false}>
        <View style={{ alignItems: 'center', borderBottomColor: '#ccc', borderBottomWidth: 1, width: '100%', paddingVertical: 40, backgroundColor: 'white' }}>
          <Image source={avatar ? { uri: avatar } : require('../assets/default-avatar.png')} style={{ width: 150, height: 150, borderRadius: 75 }} />
          {isCurrentUser && (
            <TouchableOpacity onPress={handleAvatarChange}>
              <Text style={{ color: PRIMARY_COLOR, marginTop: 10 }}>Change avatar</Text>
            </TouchableOpacity>
          )}
          <Text style={{ fontSize: 24, marginTop: 10 }}>{`${firstName} ${lastName}`}</Text>
          <Text style={{ fontSize: 18, marginTop: 5 }}>{email}</Text>
        </View>
        <FlatList 
          data={posts}
          keyExtractor={(item) => item.post.id }
          renderItem={({ item }) =><Post id={item.post.id} navigation={navigation} enableNavToDetailsScreen={true} fetchFromAPI={false} />}
          onEndReached={() => fetchPosts(page)}
          onEndReachedThreshold={0.5}
          showsVerticalScrollIndicator={false}
        />
      </ScrollView>
    </BasicLayout>
  )
}
