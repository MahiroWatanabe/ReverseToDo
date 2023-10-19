import { authInstance } from "../src/api/getUser";

const getUserData = async (url: string) => {
  try {
    const response = await authInstance.get(url);
    console.log(response);
  } catch (error) {
    console.error(error);
  }
};

const postUser = async (url: string, username: string, email: string) => {
  try {
    const response = await authInstance.post(url, { username, email });
  } catch (error) {
    console.error(error);
  }
};

export { getUserData, postUser };
