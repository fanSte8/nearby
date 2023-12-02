import { NavigationProp, useNavigation } from "@react-navigation/native"
import React from "react"
import { View, Text, StyleSheet, Image, TouchableOpacity } from "react-native"
import Ionics from "@expo/vector-icons/Ionicons"
import { SafeAreaView } from "react-native-safe-area-context"

interface PropsType {
  children: React.ReactNode
  navigation?: any
}

export const NearbyLogoLayout = ({ children, navigation }: PropsType) => {
  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.header}>
        {
          navigation && (
            <TouchableOpacity onPress={navigation.goBack} style={styles.back}>
              <Ionics name="chevron-back" color={"black"} size={32}/>
            </TouchableOpacity>
          )
        }
        <Image source={require('../assets/logo.png')} style={styles.logo}/>
        <Text style={styles.title}>Nearby</Text>
      </View>
      <View style={styles.content}>
        {children}
      </View>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  back: {
    position: 'absolute',
    top: 0,
    left: 0
  },
  container: {
    flex: 1,
    backgroundColor: '#fff',
    paddingTop: 40
  },
  header: {
    padding: 40,
    alignItems: 'center',
    marginBottom: 20
  },
  logo: {
    width: 100,
    height: 100,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold'
  },
  content: {
    flex: 1,
    alignItems: 'center',
    height: 100
  },
})
