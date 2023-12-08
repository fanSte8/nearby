import { getPosts } from "../api/posts"
import { Post, SidePanel } from "../components"
import React, { useState, useEffect, useRef } from 'react'
import { View, Text, TouchableOpacity, FlatList, StyleSheet, Modal } from 'react-native'
import { Entypo, Ionicons } from '@expo/vector-icons' 
import { SafeAreaView } from "react-native-safe-area-context"
import { PRIMARY_COLOR } from "../constants"
import DropdownPicker from 'react-native-dropdown-picker'
import * as Location from 'expo-location'
import { usePostsStore } from "../storage/usePostsStorage"
import { hasSeenNotifications as hasUnseenNotifications } from "../api/notifications"
import { shallow } from 'zustand/shallow'

export const HomeScreen = ({ navigation, route }: any) => {
  const posts = usePostsStore(store => store.posts)
  const addPosts = usePostsStore(store => store.addPosts)
  const reset = usePostsStore(store => store.reset)

  const pageSize = 20
  const [page, setPage] = useState(1)
  const [hasMorePosts, setHasMorePosts] = useState(true)
  const [isLoadingPosts, setIsLoadingPosts] = useState(false)

  const [hasNewNotifications, setHasNewNotifications] = useState(false)
  const [sortBy, setSortBy] = useState('popular')
  const [dropdownOpen, setDropdownOpen] = useState(false)
  const [dropdownItems, setDropdownItems] = useState([
    { label: 'Closest', value: 'closest' },
    { label: 'Latest', value: 'latest' },
    { label: 'Popular', value: 'popular' },
  ])

  const flatListRef = useRef<any>(null)

  useEffect(() => {
    (async () => {
      setPage(1)
      setHasMorePosts(true)
      reset()
      if (flatListRef.current) {
        flatListRef.current.scrollToOffset({ offset: 0, animated: true })
      }
      await fetchPosts(1)
    })()
  }, [sortBy])

  const [latitude, setLatitude] = useState(0)
  const [longitude, setLongitude] = useState(0)

  useEffect(() => {
    (async () => {
      const location = await Location.getCurrentPositionAsync({})
      setLatitude(location.coords.latitude)
      setLongitude(location.coords.longitude)
    })()
  }, [])

  useEffect(() => {
    (async () => {
      const res = await hasUnseenNotifications()
      setHasNewNotifications(res)
    })()
  }, [])

  const [showSidePanel, setShowSidePanel] = useState(false)
  const toggleSidePanel = () => {
    setShowSidePanel(!showSidePanel)
  }

  const fetchPosts = async (nextPage = page) => {
    if (!hasMorePosts || isLoadingPosts) {
      return
    }

    setIsLoadingPosts(true)
    const res = await getPosts(sortBy, longitude, latitude, nextPage, pageSize)
    const newPosts = res?.posts || []

    if (newPosts.length < pageSize) {
      setHasMorePosts(false)
    }
    
    addPosts(newPosts)
    setPage(nextPage + 1)
    setIsLoadingPosts(false)
  }

  useEffect(() => {
    setPage(1)
    reset()
  }, [])

  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.header}>
        <TouchableOpacity style={styles.headerIcon} onPress={toggleSidePanel}>
          <Entypo name="menu" size={28} color="white" />
        </TouchableOpacity>
        <Modal visible={showSidePanel} animationType="none" transparent={true}>
          <View style={styles.modalContainer}>
            <SidePanel onClose={toggleSidePanel} navigation={navigation} />
          </View>
        </Modal>
        <View style={styles.sortDropdownContainer}>
          <DropdownPicker
            items={dropdownItems}
            setItems={setDropdownItems}
            open={dropdownOpen}
            setOpen={setDropdownOpen}
            value={sortBy}
            setValue={setSortBy}
            multiple={false}
            style={{
              borderWidth: 0,
              backgroundColor: PRIMARY_COLOR,
              minHeight: 30,
            }}
            labelStyle={{
              backgroundColor: PRIMARY_COLOR,
              color: "white"
            }}
            textStyle={{
              textAlign: 'center',
              fontSize: 20
            }}
            showTickIcon={false}
            ArrowUpIconComponent={() => (
              <View><Ionicons name="chevron-up" color={"white"} size={16}/></View>
            )}
            ArrowDownIconComponent={() => (
              <View><Ionicons name="chevron-down" color={"white"} size={16}/></View>
            )}
          />
        </View>
        <TouchableOpacity style={styles.headerIcon} onPress={() => navigation.navigate('Notifications')}>
          <Ionicons name="notifications" size={24} color="white" />
          {hasNewNotifications && <View style={styles.notificationDot} />}
        </TouchableOpacity>
      </View>
      <FlatList
        data={posts}
        ref={flatListRef}
        keyExtractor={(item) => item.post.id }
        renderItem={({ item }) =><Post id={item.post.id} navigation={navigation} enableNavToDetailsScreen={true} />}
        onEndReached={() => fetchPosts(page)}
        onEndReachedThreshold={0.5}
      />
      <TouchableOpacity style={styles.addButton}  onPress={() => navigation.navigate('CreatePost')}>
        <Text style={styles.addButtonIcon}>+</Text>
      </TouchableOpacity>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 20,
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
    backgroundColor: PRIMARY_COLOR,
    zIndex: 2
  },
  headerIcon: {
    padding: 5,
  },
  sortDropdownContainer: {
    width: 120
  },
  notificationDot: {
    position: 'absolute',
    top: 5,
    right: 5,
    width: 10,
    height: 10,
    borderRadius: 5,
    backgroundColor: 'red',
  },
  addButton: {
    position: 'absolute',
    bottom: 20,
    right: 20,
    width: 60,
    height: 60,
    borderRadius: 30,
    backgroundColor: PRIMARY_COLOR,
    justifyContent: 'center',
    alignItems: 'center',
  },
  addButtonIcon: {
    fontSize: 40,
    color: 'white',
  },
  modalContainer: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'flex-start',
    alignItems: 'flex-start',
  },
})
