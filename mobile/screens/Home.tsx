import { View, Text, TouchableOpacity, ScrollView } from "react-native"
import { getPosts } from "../api/posts"
import { useEffect, useState } from "react"
import { Post } from "../components"
import { SafeAreaView } from "react-native-safe-area-context"

export const HomeScreen = ({ navigation }: any) => {
  const [posts, setPosts] = useState<any[]>([])
  
  useEffect(() => {
    getPosts(10, 10.1).then(response => {
      
      console.log(response.posts.map((p: any) => p.post.id))
      setPosts(response?.posts || [])
    })
  }, [])

  if (posts.length === 0) {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text>No posts found</Text>
      </View>
    )
  }

  return (
    <SafeAreaView>
      <ScrollView>
      {
        posts.map(post => <Post data={post} key={post.post.id} navigation={navigation} enableNavToDetailsScreen={true} />)
      }
      </ScrollView>
    </SafeAreaView>
  )
}
