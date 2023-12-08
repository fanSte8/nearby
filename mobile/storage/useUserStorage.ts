import AsyncStorage from '@react-native-async-storage/async-storage'
import { create } from 'zustand'
import { createJSONStorage, persist } from 'zustand/middleware'

interface User {
  id: number,
  firstName: string,
  lastName: string,
  imageUrl: string,
  email: string,
  postsRadiusKm: number,
  activated: boolean
}

interface UserState {
  user: User | null,
  isLoggedIn: boolean,
  token: string,
  setUser: (user: User) => void,
  setIsLoggedIn: (isLoggedIn: boolean) => void,
  setToken: (token: string) => void,
  reset: () => void
}

export const useUserStore = create<UserState>()(
  persist(
    (set) => ({
      user: null,
      isLoggedIn: false,
      token: '',
      setUser: (user: User) => set(() => ({ user })),
      setIsLoggedIn: (isLoggedIn: boolean) => set(() => ({ isLoggedIn })),
      setToken: (token: string) => set(() => ({ token })),
      reset: () => set(() => ({ user: null, isLoggedIn: false, token: '' }))
    }),
    {
      name: 'user-state',
      storage: createJSONStorage(() => AsyncStorage)
    }
  )
)
