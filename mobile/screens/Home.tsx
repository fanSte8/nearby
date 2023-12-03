import { View, Text, TouchableOpacity, ScrollView } from "react-native"
import { useUserStore } from "../storage/useUserStorage"
import { getPosts } from "../api/posts"
import { useEffect, useState } from "react"
import { Post } from "../components"
import { SafeAreaView } from "react-native-safe-area-context"

export const HomeScreen = () => {
  const [posts, setPosts] = useState<any[]>([])
  
  useEffect(() => {
    getPosts(10, 10.1).then(response => setPosts(response.posts))
  }, [])

  return (
    <SafeAreaView>
      <ScrollView>
      {
        posts.map(post => <Post data={post} key={post.post.id} />)
      }
      </ScrollView>
    </SafeAreaView>
  )
}
