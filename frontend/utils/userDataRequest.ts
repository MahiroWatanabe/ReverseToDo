import { authInstance } from "../src/api/getUser";

const getUserTaskData = async (url: string) => {
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

const putTaskStatus = async (url: string, status: number, id: any) => {
  try {
    const res = await authInstance.patch(url, { status, id });
    console.log(res.data);
    return res.data;
  } catch (error) {
    console.error(error);
  }
};

export { getUserTaskData, postUser, putTaskStatus };
