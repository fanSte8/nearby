import React, { useState, useEffect } from 'react'
import { View, Text, FlatList, TouchableOpacity, Image } from 'react-native'
import { getNotifications } from '../api/notifications'
import { SafeAreaView } from 'react-native-safe-area-context'
import { Ionicons } from '@expo/vector-icons'
import { BasicLayout } from '../layouts'

export const NotificationsScreen = ({ navigation }: any) => {
  const [notifications, setNotifications] = useState<any[]>([])
  const [page, setPage] = useState(1)
  const pageSize = 20

  useEffect(() => {
    fetchNotifications()
  }, [])

  const fetchNotifications = async () => {
    const newNotifications = await getNotifications(page, pageSize)

    setNotifications(prevNotifications => [...prevNotifications, ...newNotifications])
    setPage((prevPage) => prevPage + 1)
  }

  const renderNotificationItem = ({ item }: any) => {
    const { user, notifications } = item
    const notificationText = `${user.firstName} ${user.lastName} and ${
      notifications.count - 1 > 0 ? notifications.count - 1 + ' other users' : ''
    } ${notifications.type.toLowerCase()}d on your post.`

    return (
      <TouchableOpacity
        onPress={() => navigation.navigate('PostDetails', { id: notifications.postId })}
      >
        <View style={{ flexDirection: 'row', padding: 10, borderBottomWidth: 1, borderBottomColor: '#ccc' }}>
          <Image source={{ uri: user.imageUrl }} style={{ width: 50, height: 50, borderRadius: 25 }} />
          <View style={{ flex: 1, marginLeft: 10 }}>
            <Text>
              <Text style={[{fontSize: 18}, notifications.seen ? {} : { fontWeight: 'bold' }]}>{notificationText}</Text>
            </Text>
          </View>
        </View>
      </TouchableOpacity>
    )
  }

  const handleLoadMore = () => {
    fetchNotifications()
  }

  return (
    <BasicLayout navigation={navigation} title="Notifications"> 
      <View style={{ width: '100%' }}>
        <FlatList
          data={notifications}
          renderItem={renderNotificationItem}
          keyExtractor={(item, index) => index.toString()}
          onEndReached={handleLoadMore}
          onEndReachedThreshold={0.5}
        />
      </View>
    </BasicLayout>
  )
}
