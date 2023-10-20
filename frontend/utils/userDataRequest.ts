import { authInstance } from "../src/api/getUser";

const getUserData = async (url: string) => {
  try {
    const res = await authInstance.get(url);
    return res.data;
  } catch (error) {
    console.error(error);
  }
};

const postUser = async (url: string, username: string, email: string) => {
  try {
    const res = await authInstance.post(url, { username, email });
    return res.data;
  } catch (error) {
    console.error(error);
  }
};

export { getUserData, postUser };
