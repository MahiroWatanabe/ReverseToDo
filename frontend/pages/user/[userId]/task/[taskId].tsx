import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import {
  getUserTaskData,
  putTaskStatus,
} from "../../../../utils/userDataRequest";
import styles from "./task.module.scss";

const TaskComponent = () => {
  const [useTaskData, setUseTaskData] = useState<ReturnTaskData | null>(null);
  const [isChecked, setIsChecked] = useState(useTaskData?.status == 1);

  const router = useRouter();
  const currentPath: string = router.asPath;
  const pathParts = currentPath.split("/");

  const handlegetTask = (url: string) => {
    getUserTaskData(url).then((res) => {
      setUseTaskData(res);
      return res;
    });
  };

  const handleupdateStatus = (url: string, flag: number, id: any) => {
    putTaskStatus(url, flag, id).then((res) => {
      console.log(url, flag, id);
      setUseTaskData(res);
      return res;
    });
  };

  const handleUpdata = () => {};

  const handleDelete = () => {};

  const handleCheckChange = () => {
    setIsChecked((pre) => !pre);
    handleupdateStatus(`/task`, isChecked ? 1 : 0, useTaskData?.ID);
  };

  useEffect(() => {
    handlegetTask(`/task?userid=${pathParts[2]}&taskid=${pathParts[4]}`);
  }, [isChecked]);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>{useTaskData?.title}</h1>
      <p className={styles.description}>{useTaskData?.description}</p>
      <p className={styles.status}>
        Status: {isChecked ? "Active" : "Inactive"}
      </p>
      <p className={styles.deadline}>Deadline: {useTaskData?.deadline}</p>
      <button onClick={handleUpdata}>Update</button>
      <button onClick={handleDelete}>Delete</button>
      <label>
        <input
          type="checkbox"
          className={styles.checkbox}
          checked={isChecked}
          onChange={handleCheckChange}
        />
        Mark as Active
      </label>
    </div>
  );
};

export default TaskComponent;
