import React, { useEffect, useState } from 'react'
import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native'
import { formatDistance, formatTime } from '../utils'
import AntDesign from '@expo/vector-icons/AntDesign'
import FontAwesome from '@expo/vector-icons/FontAwesome'
import { PRIMARY_COLOR } from '../constants'
import { getPostById, likePost } from '../api/posts'
import { usePostsStore } from '../storage/usePostsStorage'
import { Loading } from './Loading'

export const Post = ({ id, navigation, enableNavToDetailsScreen, fetchFromAPI = false }: any) => {
  const setLiked = usePostsStore(store => store.setLiked)
  const incrementPostLikes = usePostsStore(store => store.incrementPostLikes)
  const decrementPostLikes = usePostsStore(store => store.decrementPostLikes)
  const postFromState = usePostsStore(store => store.getPostById(id))

  const [post, setPost] = useState<any>()
  const [imgWidth, setImgWidth] = useState(1)
  const [imgHeight, setImgHeight] = useState(1)
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    (
      async () => {
        let showPost
        setLoading(true)
        if (fetchFromAPI) {
          showPost = await getPostById(id)
        } else {
          showPost = postFromState
        }
        setPost(showPost)
        setLoading(false)
      })()
  }, [])

  if (!post) {
    return null
  }

  const handlePostLiked = async () => {
    setLiked(id)
    if (!post.post.liked) {
      incrementPostLikes(id)
    } else {
      decrementPostLikes(id)
    }

    setPost((state: any) => ({ 
      ...state, 
      post: { ...state.post, liked: !state.post.liked, likes: state.post.likes + (state.post.liked ? -1 : 1 ) } })
    )

    await likePost(id)
  }

  const goToPostDetails = () => {
    if (enableNavToDetailsScreen) navigation.navigate('PostDetails', { id })
  }

  const getImageDimensions = (imageUrl: string) => {
    Image.getSize(
      imageUrl,
      (width, height) => {
        setImgHeight(height)
        setImgWidth(width)
      },
      (error) => {
        console.error('Error getting image dimensions:', error)
      }
    )
  }

  if (loading) {
    return <Loading />
  }

  return (
    <View style={styles.postContainer}>
      <View style={styles.userInfoContainer}>
        <TouchableOpacity activeOpacity={1} onPress={() => navigation.navigate('Account', { id: post?.user?.id })}>
          {post?.user?.imageUrl ?(<Image
            source={{ uri: post?.user?.imageUrl }}
            style={styles.avatar}
          />) : (
            <Image
              source={require('../assets/default-avatar.png')}
              style={styles.avatar}
            />
          )}
        </TouchableOpacity>
        <View>
          <TouchableOpacity activeOpacity={1} onPress={() => navigation.navigate('Account', { id: post?.user?.id })}>
            <Text style={{fontSize: 18}}>{`${post?.user?.firstName} ${post?.user?.lastName}`}</Text>
          </TouchableOpacity>
          <View style={{flexDirection: 'row'}}>
            <Text>{formatDistance(post.post.distance)}</Text>
            <Text>{' | '}</Text>
            <Text>{formatTime(post.post.createdAt)}</Text>
          </View>
        </View>
      </View>
      <TouchableOpacity onPress={goToPostDetails} disabled={!enableNavToDetailsScreen} activeOpacity={1}>
        <Text style={{ margin: 10 }}>{post.post.description}</Text>
        <Image
          source={{ uri: post.post.imageUrl }}
          onLoad={() => getImageDimensions(post.post.imageUrl)}
          style={{
            width: '100%',
            height: 'auto',
            aspectRatio: imgWidth / imgHeight,
            resizeMode: 'contain',
          }}/>
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
    borderRadius: 20
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
