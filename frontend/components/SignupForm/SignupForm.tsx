import React from "react";
import { useState, FormEvent } from "react";
import { postUser } from "../../utils/userDataRequest";
import { useRouter } from "next/router";

const SignupForm = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const router = useRouter();

  const handleSubmit = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    // ログイン処理をここに書く
    postUser(url, username, email).then((res) => {
      router.push(`/user/${res.ID}`);
    });
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
        <button type="submit">SignUp</button>
      </form>
    </div>
  );
};

export default SignupForm;
