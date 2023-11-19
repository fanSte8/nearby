import { Text } from 'react-native'

interface PropsType {
  type: 'success' | 'warning',
  text: string
}

export const Alert = ({ text, type }: PropsType) => {
  return (
    <Text style={{ color: type === 'success' ? 'green' : 'red', alignSelf: 'center' }}>{text}</Text>
  )
}
