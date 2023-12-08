import React from 'react'
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native'
import { Entypo } from '@expo/vector-icons'
import { PRIMARY_COLOR } from '../constants'
import { useUserStore } from '../storage/useUserStorage'

export const SidePanel = ({ onClose, navigation }: any) => {
  const reset = useUserStore(store => store.reset)

  const navigateTo = (location: string) => {
    onClose()
    navigation.navigate(location)
  }

  return (
    <View style={styles.sidePanel}>
      <TouchableOpacity onPress={onClose} style={styles.closeButton}>
        <Entypo name="cross" size={24} color="black" />
      </TouchableOpacity>
      <TouchableOpacity style={styles.sidePanelButton} onPress={() => navigateTo('Account')}>
        <Text style={styles.text}>Account</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.sidePanelButton} onPress={() => navigateTo('Activate')}>
        <Text style={styles.text}>Activate</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.sidePanelButton} onPress={() => navigateTo('ChangePassword')}>
        <Text style={styles.text}>Change Password</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.sidePanelButton} onPress={() => navigateTo('ChangeRadius')}>
        <Text style={styles.text}>Change Radius</Text>
      </TouchableOpacity>
      <View style={styles.logoutContainer}>
        <TouchableOpacity style={styles.logoutButton} onPress={reset}>
          <Text style={[styles.text, { color: 'white' }]}>Log Out</Text>
        </TouchableOpacity>
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  sidePanel: {
    width: 250,
    height: '100%',
    backgroundColor: 'white',
    paddingTop: 50,
    paddingHorizontal: 20,
    position: 'absolute',
  },
  closeButton: {
    position: 'absolute',
    top: 10,
    right: 10,
  },
  sidePanelButton: {
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  logoutContainer: {
    flex: 1,
    justifyContent: 'flex-end',
    marginBottom: 20,
  },
  logoutButton: {
    paddingVertical: 10,
    backgroundColor: PRIMARY_COLOR,
    borderRadius: 5,
    alignItems: 'center',
  },
  text: {
    fontSize: 16
  }
})