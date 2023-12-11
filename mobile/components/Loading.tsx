import React from 'react'
import { View, ActivityIndicator, StyleSheet } from 'react-native'
import { PRIMARY_COLOR } from '../constants'

export const Loading = () => {
  return (
    <View style={styles.container}>
      <ActivityIndicator size="large" color={PRIMARY_COLOR} />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
})

