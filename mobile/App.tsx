import { StyleSheet, Text, View } from 'react-native';
import { LoginScreen } from './screens';
import { AuthLayout } from './layouts';

const App = () => {
  return (
    <View style={styles.container}>
      <AuthLayout>
        <LoginScreen />
      </AuthLayout>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-start',
  },
});

export default App;
