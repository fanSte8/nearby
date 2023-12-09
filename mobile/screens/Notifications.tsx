import React, { useState, useEffect } from 'react'
import { View, Text, FlatList, TouchableOpacity, Image } from 'react-native'
import { getNotifications } from '../api/notifications'
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
    
    let text = `${user.firstName} ${user.lastName} `

    switch (notifications.count) {
      case 1:
          text += 'has '
          break
      case 2:
        text += 'and 1 other user have '
        break
      default:
        text += `and ${notifications.count} other users have `
    }

    text += `${notifications.type === 'Like' ? 'liked' : 'commented on'} your post`

    return (
      <TouchableOpacity
        onPress={() => navigation.navigate('PostDetails', { id: notifications.postId })}
      >
        <View style={{ flexDirection: 'row', padding: 10, borderBottomWidth: 1, borderBottomColor: '#ccc' }}>
          <Image source={user.imageUrl ? { uri: user.imageUrl } : require('../assets/default-avatar.png')} style={{ width: 50, height: 50, borderRadius: 25 }} />
          <View style={{ flex: 1, marginLeft: 10 }}>
            <Text>
              <Text style={[{fontSize: 18}, notifications.seen ? {} : { fontWeight: 'bold' }]}>{text}</Text>
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
