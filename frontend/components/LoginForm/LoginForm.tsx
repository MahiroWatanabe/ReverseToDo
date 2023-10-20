import { useState, FormEvent } from "react";
import { getUserData } from "../../utils/userDataRequest";
import { useRouter } from "next/router";

const LoginForm = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const router = useRouter();

  const handleSubmit = (e: FormEvent<HTMLFormElement>, url: string) => {
    e.preventDefault();
    getUserData(url).then((res) => {
      router.push(`/user/${res.ID}`);
    });
  };

  return (
    <div>
      <h1>Login（GET/users?username={username}）</h1>
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
