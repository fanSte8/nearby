import AsyncStorage from '@react-native-async-storage/async-storage';
import { create } from 'zustand';
import { createJSONStorage, persist } from 'zustand/middleware';

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
  setUser: (user: User) => void,
  setIsLoggedIn: (isLoggedIn: boolean) => void,
  reset: () => void
}

export const useUserStore = create<UserState>()(
  persist(
    (set) => ({
    user: null,
    isLoggedIn: false,
    setUser: (user: User) => set(() => ({ user })),
    setIsLoggedIn: (isLoggedIn: boolean) => set(() => ({ isLoggedIn })),
    reset: () => set(() => ({ user: null, isLoggedIn: false }))
    }),
    {
      name: 'user-state',
      storage: createJSONStorage(() => AsyncStorage)
    }
  )
);
