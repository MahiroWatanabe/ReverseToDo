import { useState, FormEvent } from "react";
import { getUserData } from "../utils/userDataRequest";

const LoginForm = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");

  const handleSubmit = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    // ログイン処理をここに書く
    getUserData(url);
  };

  return (
    <div>
      <h1>Login（GET/users/&{username}）</h1>
      <form onSubmit={(e) => handleSubmit(e, `/user?username=${username}`)}>
        <label>
          Username:
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </label>
        <label>
          Email:
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </label>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default LoginForm;
