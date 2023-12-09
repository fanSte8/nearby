import React, { useEffect, useState } from "react"
import { FlatList, TextInput, TouchableOpacity, View, Text, StyleSheet, ScrollView } from "react-native"
import Ionicons from '@expo/vector-icons/Ionicons'
import { Post, Comment, Button, Input } from "../components"
import { PRIMARY_COLOR } from "../constants"
import { SafeAreaView } from "react-native-safe-area-context"
import { getComments, postComment } from "../api/posts"
import { useUserStore } from "../storage/useUserStorage"
import { usePostsStore } from "../storage/usePostsStorage"
import { BasicLayout } from "../layouts"

export const PostDetails = ({ navigation, route }: any) => {
  const pageSize = 20
  const post = route.params.id

  const user = useUserStore(store => store.user)
  const incrementPostComments = usePostsStore(store => store.incrementPostComments)

  const [comments, setComments] = useState<any[]>([])
  const [newComment, setNewComment] = useState('')
  const [page, setPage] = useState(1)
  const [hasMoreComments, setHasMoreComments] = useState(true)

  useEffect(() => {
    (async () => {
      await handleLoadMoreComments()
    })()
  }, [])

  const handleLoadMoreComments = async () => {
    if (!hasMoreComments) return
    const data = await getComments(post, page, pageSize)


    setPage(p => p + 1)
    if (!data.comments || data.comments.length < pageSize) {
      setHasMoreComments(false)
    }

    setComments([ ...comments, ...data.comments ])
  }

  const handleAddComment = async () => {
    const res = await postComment(post, newComment)
    
    if (res) {
      setNewComment('')
      incrementPostComments(post)
      setComments([{ user, comment: res.comment }, ...comments ])
    }
  }

  return (
    <BasicLayout navigation={navigation} title="Post Details">
      <ScrollView showsVerticalScrollIndicator={false}>
        <View style={styles.container}>
          <View style={{ width: '100%' }}>
            <Post id={post} navigation={navigation} enableNavToDetailsScreen={false} />
          </View>
          {user?.activated && (
            <View style={styles.addCommentContainer}>
              <Input
                placeholder="Add a comment..."
                value={newComment}
                onChangeText={(text) => setNewComment(text)}
              />
              <Button text="Post" onPress={handleAddComment} />
            </View>
          )}
          <View style={styles.commentsList}>
            <FlatList
              data={comments}
              keyExtractor={(item) => item.comment.id}
              renderItem={({ item }) => <Comment comment={item} />}
              onEndReached={() => {
                console.log('end reached')
                handleLoadMoreComments()
              }}
              onEndReachedThreshold={0.5}
              showsVerticalScrollIndicator={false}
            />
          </View>
        </View>
      </ScrollView>
    </BasicLayout>
  )
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: 'white',
    flex: 1,
    alignItems: 'center',
    width: '100%'
  },
  addCommentContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginLeft: 5,
    width: '75%'
  },
  commentInput: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    padding: 8,
    marginRight: 10,
  },
  commentButton: {
    backgroundColor: PRIMARY_COLOR,
    borderRadius: 5,
  },
  commentButtonText: {
    color: 'white',
    fontWeight: 'bold',
  },
  commentsList: {
    borderTopWidth: 1,
    borderColor: '#ccc',
    width: '100%'
  }
})