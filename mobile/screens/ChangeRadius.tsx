import { useEffect, useState } from "react"
import { NearbyLogoLayout } from "../layouts"
import { Alert, Button, LabeledInput } from "../components"
import { View, StyleSheet, Text } from "react-native"
import { PRIMARY_COLOR } from "../constants"
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view"
import { changeRadius } from "../api/users"
import { useUserStore } from "../storage/useUserStorage"

export const ChangeRadius = ({ navigation }: any) => {
  const [radius, setRadius] = useState(0)
  const [error, setError] = useState('')
  const [alert, setAlert] = useState('')

  const user = useUserStore(store => store.user)
  const setUser = useUserStore(store => store.setUser)

  useEffect(() => {
    setRadius(user?.postsRadiusKm || 0)
  }, [user])

  const handleChangeRadius = async () => {
    const { error } = await changeRadius(radius)
    if (error) {
      setAlert('')
      setError(error)
    } else {
      setUser({ ...user!, postsRadiusKm: radius })
      setError('')
      setAlert('Radius updated')
    }
  }

  return (
    <NearbyLogoLayout navigation={navigation}>
      <View style={styles.container}>
        <Text style={styles.title}>Change Radius</Text>
        <Text style={styles.text}>Change the radius of the area from which you will be able to see posts</Text>
        {
          error && <Alert type='warning' text={error} />
        }
        {
          alert && <Alert type='success' text={alert} />
        }
        <KeyboardAwareScrollView showsVerticalScrollIndicator={false}>  
          <LabeledInput
            value={String(radius)}
            onChangeText={v => setRadius(Number(v))}
            label="Radius"
            placeholder=""
            keyboardType="numeric"
          />
          <Button onPress={handleChangeRadius} text="Change Radius" />
        </KeyboardAwareScrollView>
      </View>
    </NearbyLogoLayout>
  )
}

const styles = StyleSheet.create({
  title: {
    padding: 10,
    textAlign: "center",
    fontSize: 18
  },
  text: {
    padding: 10,
    textAlign: "center",
  },
  container: {
    backgroundColor: '#ffffff',
    width: "100%",
    height: "100%",
    padding: 10
  },
  link: {
    alignSelf: 'center',
    paddingTop: 10
  },
  linkText: {
    color: PRIMARY_COLOR
  }
})
