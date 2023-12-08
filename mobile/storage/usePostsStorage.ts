import AsyncStorage from '@react-native-async-storage/async-storage'
import { create } from 'zustand'
import { createJSONStorage, persist } from 'zustand/middleware'

interface PostsState {
  posts: any[],
  addPosts: (posts: any[]) => void,
  setLiked: (postId: number) => void,
  incrementPostLikes: (postId: number) => void
  incrementPostComments: (postId: number) => void
  decrementPostLikes: (postId: number) => void
  decrementPostComments: (postId: number) => void
  getPostById: (postId: number) => any
  reset: () => void
}

export const usePostsStore = create<PostsState>((set, get) => ({
  posts: [],
  addPosts: (posts) => set((state) => ({ posts: [...state.posts, ...posts] })),
  setLiked: (postId) => 
    set((state) => ({
      posts: state.posts.map((post) =>
        post.post.id === postId ? { ...post, post: { ...post.post, liked: !post.post.liked  } } : post
      ),
    })),
  incrementPostLikes: (postId) =>
    set((state) => ({
      posts: state.posts.map((post) =>
        post.post.id === postId ? { ...post, post: { ...post.post, likes: post.post.likes + 1 } } : post
      ),
    })),
  incrementPostComments: (postId) =>
    set((state) => ({
      posts: state.posts.map((post) =>
        post.post.id === postId ? { ...post, post: { ...post.post, comments: post.post.comments + 1 } } : post
      ),
    })),
  decrementPostLikes: (postId) =>
    set((state) => ({
      posts: state.posts.map((post) =>
        post.post.id === postId ? { ...post, post: { ...post.post, likes: post.post.likes - 1 } } : post
      ),
    })),
  decrementPostComments: (postId) =>
    set((state) => ({
      posts: state.posts.map((post) =>
        post.post.id === postId ? { ...post, post: { ...post.post, comments: post.post.comments - 1 } } : post
      ),
    })),
  getPostById: (postId) => {
    const currentState = get()
    return currentState.posts.find((post) => post.post.id === postId)
  },
  reset: () => set({ posts: [] }),
}))
