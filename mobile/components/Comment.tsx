import { View, Text, Image, StyleSheet } from "react-native"
import { formatTime } from "../utils"

export const Comment = ({ comment }: any) => {
  const { user: { firstName, lastName, avatarUrl }, comment: { text, createdAt } } = comment

  return (
    <View style={styles.commentContainer}>
      {avatarUrl ?(<Image
          source={{ uri: avatarUrl }}
          style={styles.commentAvatar}
        />) : (
          <Image
            source={require('../assets/default-avatar.png')}
            style={styles.commentAvatar}
          />
        )}
      <View style={styles.commentContent}>
        <Text style={{fontSize: 14, color: '#555'}}>{`${firstName} ${lastName}`}</Text>
        <Text style={{fontSize: 16}}>{text}</Text>
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  commentContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingBottom: 10,
    borderBottomWidth: 1,
    borderColor: '#ccc'
  },
  commentAvatar: {
    width: 40,
    height: 40,
    borderRadius: 20,
    marginRight: 10,
  },
  commentContent: {
    flex: 1,
  },
  commentDate: {
    color: '#888',
    fontSize: 12,
  },
})