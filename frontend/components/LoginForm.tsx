import { useState, FormEvent } from "react";

import { authInstance } from "../src/api/getUser";

const LoginForm = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");

  const getUserData = async (url: string) => {
    try {
      const response = await authInstance.get(url);
    } catch (error) {
      console.error(error);
    }
  };
  const postUser = async (url: string) => {
    try {
      const response = await authInstance.post(url, { username, email });
    } catch (error) {
      console.error(error);
    }
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    // ログイン処理をここに書く
    getUserData(url);
    console.log(`Username: ${username}, Email: ${email}`);
  };

  const handleSubmit2 = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    // ログイン処理をここに書く
    postUser(url);
    console.log(`Username: ${username}, Email: ${email}`);
  };
  return (
    <div>
      <h1>Login（GET/users/&{username}）</h1>
      <form onSubmit={(e) => handleSubmit(e, `/users/${username}`)}>
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

      <h1>ユーザ作成（POST/）</h1>
      <form onSubmit={(e) => handleSubmit2(e, "/users")}>
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
