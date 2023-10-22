import { useState, useEffect } from "react";
import styles from "./UserTask.module.scss";
import { getUserTaskData } from "../../utils/userDataRequest";
import TodoItem from "../TodoItem/TodoItem";

type Props = {
  id: string;
};

const UserTask = ({ id }: Props) => {
  const [useUserData, setUseUserData] = useState<ReturnUserData | null>(null);
  const [useTaskData, setUseTaskData] = useState<ReturnTaskData[] | null>(null);

  const handleSubmit = (url: string) => {
    getUserTaskData(url).then((res) => {
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
