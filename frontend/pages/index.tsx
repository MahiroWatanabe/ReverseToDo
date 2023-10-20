import type { NextPage } from "next";
import { useRouter } from "next/router";
import styles from "../styles/Home.module.css";
import LoginForm from "../components/LoginForm/LoginForm";

const Home: NextPage = () => {
  const router = useRouter();

  const handleSignupClick = () => {
    router.push("/signup");
  };

  return (
    <div>
      <LoginForm />
      <p>
        サインアップは
        <span className={styles.blueText} onClick={handleSignupClick}>
          こちら
        </span>
      </p>
    </div>
  );
};

export default Home;
