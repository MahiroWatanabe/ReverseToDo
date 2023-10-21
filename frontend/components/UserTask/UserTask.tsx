import { useState, FormEvent, useEffect } from "react";
import styles from "./UserTask.module.scss";
import { getUserData } from "../../utils/userDataRequest";
import { useRouter } from "next/router";
import TodoItem from "../TodoItem/TodoItem";

type Props = {
  id: string;
};

type ReturnUserData = {
  CreatedAt: string;
  DeletedAt: string;
  ID: number;
  Tasks: [];
  UpdatedAt: string;
  email: string;
  username: string;
};

type ReturnTaskData = {
  CreatedAt: string;
  DeletedAt: string | null;
  ID: number;
  UpdatedAt: string;
  assignid: number;
  createid: number;
  deadline: string;
  description: string;
  status: number;
  title: string;
};

const UserTask = ({ id }: Props) => {
  const [useUserData, setUseUserData] = useState<ReturnUserData | null>(null);
  const [useTaskData, setUseTaskData] = useState<ReturnTaskData[] | null>(null);

  const handleSubmit = (url: string) => {
    getUserData(url).then((res) => {
      setUseUserData(res?.userdata);
      setUseTaskData(res?.taskdata);
      return res;
    });
  };

  useEffect(() => {
    handleSubmit(`/user?id=${id}`);
  }, []);

  return (
    <div className={styles["user-task-container"]}>
      <h1>Login（GET/users?username={useUserData?.username}）</h1>

      {useTaskData?.map((data) => (
        <TodoItem key={data.ID} item={data} />
      ))}
    </div>
  );
};

export default UserTask;
