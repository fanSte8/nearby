import React, { useEffect, useState } from 'react'
import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native'
import { formatDistance, formatTime } from '../utils'
import AntDesign from '@expo/vector-icons/AntDesign'
import FontAwesome from '@expo/vector-icons/FontAwesome'
import { PRIMARY_COLOR } from '../constants'
import { likePost } from '../api/posts'

export const Post = ({ data, navigation, enableNavToDetailsScreen }: any) => {
  const [liked, setLiked] = useState(false)
  const [likes, setLikes] = useState(0)

  useEffect(() => {
    setLiked(data.post.liked)
    setLikes(data.post.likes)
  }, [])

  const {
    user: {
      firstName,
      lastName,
      imageUrl: avatarUrl
    },
    post: {
      description,
      imageUrl,
      distance,
      comments,
      createdAt
    }
  } = data

  const handlePostLiked = async () => {
    const success = likePost(data.post.id)
    
    if (!success) return

    if (liked) {
      setLikes(likes => likes - 1)
    } else {
      setLikes(likes => likes + 1)
    }
    setLiked(liked => !liked)
  }

  const goToPostDetails = () => {
    if (enableNavToDetailsScreen) navigation.navigate('PostDetails', { data })
  }

  return (
    <View style={styles.postContainer}>
      <View style={styles.userInfoContainer}>
        {avatarUrl ?(<Image
          source={{ uri: avatarUrl }}
          style={styles.avatar}
        />) : (
          <Image
            source={require('../assets/default-avatar.png')}
            style={styles.avatar}
          />
        )}
        <View>
          <Text style={{fontSize: 18}}>{`${firstName} ${lastName}`}</Text>
          <View style={{flexDirection: 'row'}}>
            <Text>{formatDistance(distance)}</Text>
            <Text>{' | '}</Text>
            <Text>{formatTime(createdAt)}</Text>
          </View>
        </View>
      </View>
      <TouchableOpacity onPress={goToPostDetails} disabled={!enableNavToDetailsScreen} activeOpacity={1}>
        <Text style={{ margin: 10 }}>{description}</Text>
        <Image source={{ uri: imageUrl }} style={styles.postImage} />
      </TouchableOpacity>
      <View style={styles.actionButtonsContainer}>
       <TouchableOpacity style={styles.button} onPress={handlePostLiked}>
          <Text style={liked ? styles.buttonTextWithPrimaryColor : styles.buttonText}>{likes}</Text>
          <AntDesign name='like2' size={18} color={liked ? PRIMARY_COLOR : '#000'} />
        </TouchableOpacity>
        <TouchableOpacity style={styles.button} onPress={goToPostDetails} disabled={!enableNavToDetailsScreen} activeOpacity={1}>
          <Text style={styles.buttonText}>{comments}</Text>
          <FontAwesome name='comment-o' size={18} />
        </TouchableOpacity>
      </View>
    </View>
  )
}


const styles = StyleSheet.create({
  postContainer: {
    backgroundColor: '#fff',
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  userInfoContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 10,
    marginTop: 10,
  },
  avatar: {
    width: 40,
    height: 40,
    marginRight: 10,
  },
  postImage: {
    width: '100%',
    height: 200,
  },
  actionButtonsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center'
  },
  button: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginVertical: 3
  },
  buttonText: {
    fontSize: 18,
    marginHorizontal: 5,
  },
  buttonTextWithPrimaryColor: {
    color: PRIMARY_COLOR,
    fontSize: 18,
    marginHorizontal: 5,
  },
})
