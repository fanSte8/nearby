import React, { useEffect, useState } from 'react'
import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native'
import { formatDistance, formatTime } from '../utils'
import AntDesign from '@expo/vector-icons/AntDesign'
import FontAwesome from '@expo/vector-icons/FontAwesome'
import { PRIMARY_COLOR } from '../constants'
import { likePost } from '../api/posts'
import { usePostsStore } from '../storage/usePostsStorage'

export const Post = ({ id, navigation, enableNavToDetailsScreen }: any) => {
  const setLiked = usePostsStore(store => store.setLiked)
  const incrementPostLikes = usePostsStore(store => store.incrementPostLikes)
  const decrementPostLikes = usePostsStore(store => store.decrementPostLikes)
  const post = usePostsStore(store => store.getPostById(id))

  if (!post) {
    return null
  }

  const [updateFlag, setUpdateFlag] = useState(0)

  const handlePostLiked = async () => {
    setLiked(id)
    if (!post.post.liked) {
      incrementPostLikes(id)
    } else {
      decrementPostLikes(id)
    }

    await likePost(id)
  }

  const goToPostDetails = () => {
    if (enableNavToDetailsScreen) navigation.navigate('PostDetails', { id })
  }

  return (
    <View style={styles.postContainer}>
      <View style={styles.userInfoContainer}>
        {post.user.avatarUrl ?(<Image
          source={{ uri: post.user.avatarUrl }}
          style={styles.avatar}
        />) : (
          <Image
            source={require('../assets/default-avatar.png')}
            style={styles.avatar}
          />
        )}
        <View>
          <Text style={{fontSize: 18}}>{`${post.user.firstName} ${post.user.lastName}`}</Text>
          <View style={{flexDirection: 'row'}}>
            <Text>{formatDistance(post.post.distance)}</Text>
            <Text>{' | '}</Text>
            <Text>{formatTime(post.post.createdAt)}</Text>
          </View>
        </View>
      </View>
      <TouchableOpacity onPress={goToPostDetails} disabled={!enableNavToDetailsScreen} activeOpacity={1}>
        <Text style={{ margin: 10 }}>{post.post.description}</Text>
        <Image source={{ uri: post.post.imageUrl }} style={styles.postImage} />
      </TouchableOpacity>
      <View style={styles.actionButtonsContainer}>
       <TouchableOpacity style={styles.button} onPress={handlePostLiked}>
          <Text style={post.post.liked ? styles.buttonTextWithPrimaryColor : styles.buttonText}>{post.post.likes}</Text>
          <AntDesign name='like2' size={18} color={post.post.liked ? PRIMARY_COLOR : '#000'} />
        </TouchableOpacity>
        <TouchableOpacity style={styles.button} onPress={goToPostDetails} disabled={!enableNavToDetailsScreen} activeOpacity={1}>
          <Text style={styles.buttonText}>{post.post.comments}</Text>
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
