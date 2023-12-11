import { NavigationProp, useNavigation } from "@react-navigation/native"
import React from "react"
import { View, Text, StyleSheet, Image, TouchableOpacity } from "react-native"
import Ionics from "@expo/vector-icons/Ionicons"
import { SafeAreaView } from "react-native-safe-area-context"
import { PRIMARY_COLOR } from "../constants"

interface PropsType {
  children: React.ReactNode
  navigation?: any,
  title: string
}

export const BasicLayout = ({ children, navigation, title }: PropsType) => {
  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.header}>
        {
          navigation && (
            <TouchableOpacity onPress={navigation.goBack} style={styles.back}>
              <Ionics name="chevron-back" color={"white"} size={32}/>
            </TouchableOpacity>
          )
        }
        <Text style={styles.title}>{title}</Text>
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
    left: 0
  },
  container: {
    flex: 1
  },
  header: {
    backgroundColor: PRIMARY_COLOR,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 10,
  },
  content: {
    flex: 1,
    alignItems: 'center',
    height: 100
  },
  title: {
    color: 'white',
    fontSize:  24
  }
})
