import React from "react";
import { View, Text, StyleSheet, Image } from "react-native";

export const AuthLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <View style={styles.container}>
      <View style={styles.header}>
        <Image source={require('../assets/logo.png')} style={styles.logo}/>
        <Text style={styles.title}>Nearby</Text>
      </View>
      <View style={styles.content}>
        {children}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 40,
  },
  header: {
    padding: 40,
    alignItems: 'center',
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
    justifyContent: 'center'
  },
});
