import { useState, FormEvent, useEffect } from "react";
import styles from "./UserTask.module.scss";
import { getUserData } from "../../utils/userDataRequest";
import { useRouter } from "next/router";

type Props = {
  id: string;
};

type ReturnData = {
  CreatedAt: string;
  DeletedAt: string;
  ID: number;
  Tasks: [];
  UpdatedAt: string;
  email: string;
  username: string;
};

const UserTask = ({ id }: Props) => {
  const [useData, setUserdata] = useState<ReturnData | null>(null);

  const handleSubmit = (url: string) => {
    getUserData(url).then((res) => {
      setUserdata(res);
      return res;
    });
  };

  useEffect(() => {
    handleSubmit(`/user?id=${id}`);
  }, []);

  console.log(useData);

  return (
    <div className={styles["user-task-container"]}>
      <h1>Login（GET/users?username={useData?.username}）</h1>
    </div>
  );
};

export default UserTask;
