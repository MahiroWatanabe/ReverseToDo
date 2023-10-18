import type { NextPage } from 'next'
import LoginForm from '../components/LoginForm';

const Home: NextPage = () => {
  // ログイン画面の作成から行う
  return (
    <LoginForm />
  )
}

export default Home
