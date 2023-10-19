import React from "react";
import { useState, FormEvent } from "react";
import { postUser } from "../utils/userDataRequest";

const SignupForm = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");

  const handleSubmit = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    // ログイン処理をここに書く
    postUser(url, username, email);
    console.log(`Username: ${username}, Email: ${email}`);
  };

  return (
    <div>
      <h1>ユーザ作成（POST/）</h1>
      <form onSubmit={(e) => handleSubmit(e, "/user")}>
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

export default SignupForm;
